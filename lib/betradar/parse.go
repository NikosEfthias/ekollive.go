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
)

var old = ""

func Parse(c chan models.BetradarLiveOdds) {
	con := Connect(*lib.BetradarURL)
	Login(con)
	scanner := GetBufferReader(con)
	var mainTag bytes.Buffer
	var start, flush bool
	var limiter = make(chan bool, *lib.J)
	if *lib.BAR {
		go func() {
			for {
				time.Sleep(time.Millisecond * 500)
				fmt.Printf("\rlimiter (%d)=> %d", len(limiter), runtime.NumGoroutine())
			}
		}()
	}
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
			go func(res models.BetradarLiveOdds) {
				//lib.PrintProgress(runtime.NumGoroutine(), '.')
				limiter <- true
				if res.Status == nil {
					return
				}
				if !*lib.Testing {
					go controllers.UpsertMatches(res.Match, limiter)
				}
			}(res)
			mainTag.Reset()
			flush = false
		}
	}
}
