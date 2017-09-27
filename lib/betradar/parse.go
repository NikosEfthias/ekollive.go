package betradar

import (
	"bytes"
	"encoding/xml"
	"strings"
	"../../controllers"
	"../../lib"
	"../../models"
	"fmt"
	"time"
	"runtime"
	"bufio"
)

var limiter chan bool

func init() {
	limiter = make(chan bool, *lib.J)
	if *lib.BAR {
		go func() {
			for {
				time.Sleep(time.Millisecond * 500)
				fmt.Printf("\rlimiter (%d)=> goroutinesNum(%d)", len(limiter), runtime.NumGoroutine())
			}
		}()
	}
}
func Parse(c chan models.BetradarLiveOdds) {
	con := Connect(*lib.ProxyURL)
	scanner := bufio.NewScanner(con)
	var mainTag bytes.Buffer
	var start, flush bool
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if !strings.HasSuffix(line, "/>") && strings.HasPrefix(line, "<BetradarLiveOdds") {
			//start
			start = true
			mainTag.WriteString(line)
			continue
		} else if strings.HasPrefix(line, "</BetradarLiveOdds") {
			mainTag.WriteString(line)
			start = false
			flush = true
		}
		if start {
			mainTag.WriteString(line)
		}
		if flush {
			var res = models.BetradarLiveOdds{}
			xml.Unmarshal(mainTag.Bytes(), &res)
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
					controllers.UpsertMatches(res.Match, limiter)
				}
			}(res)
			mainTag.Reset()
			flush = false
		}
	}
}
