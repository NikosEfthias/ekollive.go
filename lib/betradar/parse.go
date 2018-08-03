package betradar

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
	"github.com/mugsoft/tools/bytesize"
)

// var limiter chan bool
var dataChan chan []byte

// var marshalDone chan bool

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
	// marshalDone = make(chan bool)
}

var consuming bool
var ReadForTheFirstTime bool = false

func Parse() {
	con := Connect(*lib.ProxyURL)
	//	var marshalDone = make(chan bool)
	go sendping(con)
	go func() {
		var erroredNum int
		var goodNum int
		scanner := bufio.NewScanner(con)
		scanner.Split(bufio.ScanLines)
		scanner.Buffer(make([]byte, bytesize.MB*500), int(bytesize.GB))
		// var length int
		// const bufsize = 1024
		// var buf = make([]byte, bufsize)
		// var remainingOctets int
		// var data = make([]byte, 0, 1<<10)
		for scanner.Scan() {
			// var meta = make([]byte, 4)
			// <-marshalDone
			// data := make([]byte, 5000)
			// n, err := con.Read(data)
			// if nil != err {
			// 	pp.Println(err.Error())
			// 	break
			// } else if n < 4 {
			// 	break
			// }
			// if n == 1 {
			// 	pp.Println("n==1")
			// 	pp.Println(data[:n])
			// }
			// pp.Println(meta)
			// data = data[:0]
			// length = int(tools.LE2Int(meta))
			// remainingOctets = length
			// consuming = true
			// for remainingOctets > 0 {
			//
			// 	if remainingOctets < len(buf) {
			// 		buf = buf[:remainingOctets]
			// 	} else if remainingOctets > bufsize {
			// 		buf = make([]byte, bufsize)
			// 	}
			// 	n, err = con.Read(buf)
			// 	if nil != err {
			// 		pp.Println(err.Error())
			// 		break
			// 	}
			// 	fmt.Println(n, len(buf[:n]), remainingOctets)
			// 	remainingOctets -= n
			// 	data = append(data, buf[:n]...)
			// 	time.Sleep(time.Millisecond * 10)
			// }
			// var dta = scanner.Bytes()
			txt := scanner.Text()
			if len(txt) == 0 {
				fmt.Fprintln(os.Stderr, "0 data")
				continue
			}
			var data = map[string]interface{}{}
			err := json.Unmarshal(scanner.Bytes(), &data)
			// pp.Println(data)
			if nil != err {
				erroredNum++
				fmt.Fprintln(os.Stderr, err)
			} else {
				goodNum++
			}
			fmt.Printf("\033cgood => %d;bad => %d; %.2f%%", goodNum, erroredNum, (float64(erroredNum)*100.0)/float64(goodNum))
			// dataChan <- dta //data[:length]
		}
		log.Println("\nbetconstruct connection was interrrupted restarting", scanner.Err())
		os.Exit(1)
	}()
	// marshalDone <- true //initially allow the reador to read
	var value *models.BetconstructData
	var count = 0
	for {
		select {
		case a := <-dataChan:
			ReadForTheFirstTime = true
			pp.Println(string(a))
			value = new(models.BetconstructData)
			err := json.Unmarshal([]byte(a), value)
			if nil != err {
				//TODO[x]:  write this to a file
				if count <= 1 {
					count++
					// marshalDone <- true
					continue
				}
				f, _ := os.OpenFile("erroredJSN.json", os.O_CREATE|os.O_WRONLY, 0x755)
				f.Write([]byte(a))
				f.Close()
				pp.Println("error parsing the json from the parser", err)
				os.Exit(1)
				// marshalDone <- true
				continue
			} else if nil == value || nil == value.Command {
				//TODO:  handle this
				pp.Println("value is nil", value, string(a))
				// marshalDone <- true
				continue
			}
			// marshalDone <- true
			// case <-time.After(time.Minute * 2):
			// 	// if consuming {
			// 	// 	pp.Println("consuming")
			// 	// 	continue
			// 	// } else
			// 	if !ReadForTheFirstTime {
			// 		pp.Println("waiting for the initial message")
			// 		continue
			// 	}
			// 	fmt.Println("\n\n\n\x1B[31m", "no data for 2 minutes restarting", "\x1B[0m")
			// 	os.Exit(0)
		}
		// fmt.Print("\r", "\033c\x1B[32m", *value.Command, "  ", time.Now().Format("03:04:05"), "\x1B[0m\r")
		if nil != value.Error {
			pp.Println(value)
		}
	}
}
func sendping(con net.Conn) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		<-ticker.C
		_, err := con.Write([]byte("ping"))
		if nil != err {
			pp.Println("sock error")
			os.Exit(1)
		}
	}
}
