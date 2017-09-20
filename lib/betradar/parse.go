package betradar

import (
	"../../conf"
	"../../models"
	_ "../../controllers"
	"os"
	"bytes"
	"strings"
	"runtime"
	"encoding/xml"
	"time"
)

var old = ""

func Parse(c chan models.BetradarLiveOdds) {
	con := Connect(conf.Conf["betradar-"+os.Getenv("ekolEnv")+"-url"])
	Login(con)
	scanner := GetBufferReader(con)
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

			go func(res models.BetradarLiveOdds) {
			check:
			//fmt.Print("\x1B[2K\r", runtime.NumGoroutine())
			//lib.PrintProgress(runtime.NumGoroutine(), '.')
				if runtime.NumGoroutine() < 50 {
					//go controllers.UpsertMatches(res.Match)
				} else {
					time.Sleep(time.Millisecond * 300)
					goto check
				}
			}(res)
			mainTag.Reset()
			flush = false
		}
	}
}
