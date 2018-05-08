package betradar

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"../../controllers"
	"../../lib"
	"../../models"
	"../db"
	"../websocketops"
	"../store/oddids"
)

var limiter chan bool
var mainTag bytes.Buffer
var start, flush bool
var data chan bool

func init() {
	limiter = make(chan bool, *lib.J)
	if *lib.BAR {
		go func() {
			for {
				time.Sleep(time.Millisecond * 10)
				fmt.Printf("\rlimiter (%d)=> goroutinesNum(%d)=> connectedClients(%d)", len(limiter), runtime.NumGoroutine(), len(websocketops.SocketList))
			}
		}()
	}
	data = make(chan bool)
}

func Parse(c chan models.BetradarLiveOdds) {
	con := Connect(*lib.ProxyURL)
	scanner := bufio.NewScanner(con)

	for {
		go func() { data <- scanner.Scan() }()
		select {
		case a := <-data:
			if !a {
				time.Sleep(time.Second * 5)
				fmt.Println("\n\n\nconnection dropped reconnecting")
				db.DB.DB().Exec("update matches set betstatus='stopped' where betstatus='started'")
				db.DB.DB().Exec("update odds set active='0' WHERE active='1'")
				os.Exit(0)
			}
		case <-time.After(time.Second * 50):
			fmt.Println("\n\n\n\x1B[31m", "no data for 50 seconds restarting", "\x1B[0m")
			db.DB.DB().Exec("update matches set betstatus='stopped' where betstatus='started'")
			db.DB.DB().Exec("update odds set active='0' WHERE active='1'")
			os.Exit(0)
		}
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasSuffix(line, "/>") && strings.HasPrefix(line, "<BetradarLiveOdds") {
			//start
			//fmt.Println(line)
			start = true
			mainTag.WriteString(line)
			continue
		} else if strings.HasSuffix(line, "/>") && strings.HasPrefix(line, "<BetradarLiveOdds") {
			mainTag.WriteString(line)
			start = false
			flush = true
		} else if start && strings.HasPrefix(line, "</BetradarLiveOdds") {
			//fmt.Println(line)
			mainTag.WriteString(line)
			start = false
			flush = true
		}
		if start {
			mainTag.WriteString(line)
		}
		if flush {
			var res = models.BetradarLiveOdds{}

			err := xml.Unmarshal(mainTag.Bytes(), &res)
			if nil != err {
				f, er := os.Create("errored.tags.dump.xml")
				fmt.Println(er)
				f.Write(mainTag.Bytes())
				f.Close()
				fmt.Println("\x1B[31mXMLParseError", err, "\x1B[0m")
				mainTag.Reset()
				flush = false
				continue
			}
			for _, match := range res.Match {
				for _,odd :=range match.Odds {
					if nil != odd.OddTypeId {
						err := oddids.SetById(&odd)
						if nil != err {
							fmt.Println(err)
						}
					}
				}
			}
			go func(res models.BetradarLiveOdds) {
				var retried = false
			retry:
				select {
				case c <- res:
				case <-time.After(time.Millisecond * 10):
					if retried {
						return
					}
					retried = true
					goto retry
				}
			}(res)
			if *lib.DumpTags {
				fmt.Println("\n\n", mainTag.String())
			}
			time.Sleep(time.Millisecond * 50)
			if res.Status != nil && *res.Status == "alive" {
				goto alive
			}
			go func(res models.BetradarLiveOdds) {
				limiter <- true
				if res.Status == nil {
					return
				}
				controllers.UpsertMatches(res.Match, limiter, res)
			}(res)
		alive:
			mainTag.Reset()
			flush = false
		}
	}
}
