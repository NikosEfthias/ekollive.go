package betradar

import (
	"bytes"
	"../../lib"
	"../../models"
	"../db"
	"fmt"
	"time"
	"runtime"
	"bufio"
	"os"
	"strings"
	"encoding/xml"
	"../../controllers"
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
				time.Sleep(time.Millisecond * 100)
				fmt.Printf("\rlimiter (%d)=> goroutinesNum(%d)", len(limiter), runtime.NumGoroutine())
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
			fmt.Println("\n\n\n\x1B[31m", "no data for 50 seconds restarting", "\x1B[0m\n\n\n")
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

			select {
			case c <- res:
			default:
			}
			if *lib.DumpTags {

				fmt.Println("\n\n", mainTag.String(), "\n\n")
			}

			go func(res models.BetradarLiveOdds) {
				limiter <- true
				if res.Status == nil {
					return
				}

				if !*lib.Testing {
					controllers.UpsertMatches(res.Match, limiter, res)
				}
			}(res)
			mainTag.Reset()
			flush = false
		}
	}
}
