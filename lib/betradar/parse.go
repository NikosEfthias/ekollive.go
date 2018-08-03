package betradar

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
	"github.com/mugsoft/tools"
)

// var limiter chan bool
var dataChan chan []byte
var marshalDone chan bool

func init() {
	// limiter = make(chan bool, *lib.J)
	// if *lib.BAR {
	// 	go func() {
	// 		for {
	// 			time.Sleep(time.Millisecond * 10)
	// 			fmt.Printf("\rlimiter (%d)=> goroutinesNum(%d)", len(limiter), runtime.NumGoroutine())
	// 		}
	// 	}()
	// }
	dataChan = make(chan []byte)
	marshalDone = make(chan bool)
}

var consuming bool

func Parse() {
	con := Connect(*lib.ProxyURL)
	//	var marshalDone = make(chan bool)
	go func() {
		var meta = make([]byte, 4)
		var length int
		var data []byte
		const bufsize = 100
		var buf = make([]byte, bufsize)
		var remainingOctets int
		for {
			<-marshalDone
			n, err := con.Read(meta)
			if nil != err {
				pp.Println(err.Error())
				break
			} else if n < 4 {
				break
			}
			length = int(tools.LE2Int(meta))
			remainingOctets = length
			for remainingOctets > 0 {
				consuming = true
				if remainingOctets < len(buf) {
					n, err = con.Read(buf[:remainingOctets])
				} else {
					n, err = con.Read(buf)
				}
				if nil != err {
					pp.Println(err.Error())
					break
				}
				remainingOctets -= n
				data = append(data, buf[:n]...)
			}
			dataChan <- data
			data = data[:0]
			consuming = false
		}
		consuming = false
		log.Println("\nbetconstruct connection was interrrupted restarting")
		os.Exit(1)
	}()
	marshalDone <- true //initially allow the reador to read
	var value *models.BetconstructData
	for {
		select {
		case a := <-dataChan:
			// pp.Println(string(a))
			value = new(models.BetconstructData)
			err := json.Unmarshal(a, value)
			marshalDone <- true
			if nil != err {
				//TODO[x]:  write this to a file
				f, _ := os.OpenFile("erroredJSN.json", os.O_CREATE|os.O_WRONLY, 0x755)
				f.Write(a)
				f.Close()
				pp.Println("error parsing the json from the parser", err)
				continue
			} else if nil == value || nil == value.Command {
				//TODO:  handle this
				pp.Println("value is nil", value, string(a))
				continue
			}
		case <-time.After(time.Minute):
			if consuming {
				pp.Println("consuming")
				continue
			}
			fmt.Println("\n\n\n\x1B[31m", "no data for 11 seconds restarting", "\x1B[0m")
			os.Exit(0)
		}
		fmt.Print("\r", "\x1B[32m", *value.Command, "\x1B[0m\r")
		if nil != value.Error {
			pp.Println(value)
		}
	}
}
