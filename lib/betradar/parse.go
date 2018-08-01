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
}

var writeit = make(chan bool)

func Parse() {
	con := Connect(*lib.DirectURL)
	//	var marshalDone = make(chan bool)
	go func() {
		for {
			var totalData []byte
			var meta = make([]byte, 4)
			n, err := con.Read(meta)
			var length = int(tools.LE2Int(meta))
			var remainsToConsume = length
			if nil != err {
				log.Fatalln(err)
			} else if n < 4 {
				fmt.Println("Erroorrr they sent less bytes ")
				continue
			}

		readMore:
			var data = make([]byte, remainsToConsume)
			n, _ = con.Read(data)
			remainsToConsume -= n
			totalData = append(totalData, data[:n]...)
			pp.Println(remainsToConsume, "<===>", length)
			if 0 < remainsToConsume {
				goto readMore
			}
			dataChan <- totalData[:length]
		}
		time.Sleep(time.Second * 11)
		log.Println("\nbetconstruct connection was interrrupted restarting")
		os.Exit(1)
	}()
	// go func() {
	// 	var header1 = make([]byte, 1)
	// 	var header3 = make([]byte, 3)
	// 	var _length = make([]byte, 4)
	// 	var totalData []byte
	// 	const bufsize = bytesize.MB // bytesize.KB * 100
	// 	var data []byte = make([]byte, bufsize)
	// 	var totalRead int
	// 	var length int
	// 	for {
	// 		pp.Println("locked")
	//	// 		<-marshalDone
	// 		pp.Println("unlocked")
	// 		//TODO:  handle errors in this scope
	// 		totalRead = 0
	// 		totalData = totalData[:0]
	// 		con.Read(header1)
	// 		if string(header1) != "\r" {
	// 			go func() {
	//	// 				marshalDone <- true
	// 			}()
	// 			continue
	// 		}
	// 		con.Read(header3)
	// 		if string(header3) != "\n\r\n" {
	// 			go func() {
	//	// 				marshalDone <- true
	// 			}()
	// 			continue
	// 		}
	// 		n, _ := con.Read(_length)
	// 		if n < 4 {
	// 			go func() {
	//	// 				marshalDone <- true
	// 			}()
	// 			continue
	// 		}
	// 		length = int(tools.LE2Int(_length))
	// 	readMore:
	// 		n, _ = con.Read(data)
	// 		totalRead += n
	// 		totalData = append(totalData, data[:n]...)
	// 		pp.Println(length, totalRead)
	// 		if totalRead < length {
	// 			pp.Println("goto")
	// 			goto readMore
	// 		}
	// 		if totalRead > length {
	// 			totalRead = 0
	// 			totalData = totalData[:0]
	// 			pp.Println("read more")
	// 			go func() {
	//	// 				marshalDone <- true
	// 			}()
	// 			continue
	// 		}
	// 		dataChan <- totalData[:totalRead]
	// 	}
	// }()
	var value *models.BetconstructData
	for {
		select {
		case a := <-dataChan:
			// pp.Println(string(a))
			value = new(models.BetconstructData)
			err := json.Unmarshal(a, value)
			//			marshalDone <- true
			if nil != err {
				//TODO:  write this to a file
				// pp.Println(string(a))
				f, _ := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0x755)
				f.Write(a)
				f.Close()
				pp.Println("error parsing the json from the parser", err)
				continue
			} else if nil == value || nil == value.Command {
				//TODO:  handle this
				pp.Println("value is nil")
				continue
			}
		case <-time.After(time.Second * 11):
			fmt.Println("\n\n\n\x1B[31m", "no data for 11 seconds restarting", "\x1B[0m")
			os.Exit(0)
		}
		pp.Println(value.Command)
		if nil != value.Error {
			pp.Println(value)
		}
		// go func(res *models.BetradarLiveOdds) {
		// 	var retried = false
		// retry:
		// 	select {
		// 	case c <- res:
		// 	case <-time.After(time.Millisecond * 10):
		// 		if retried {
		// 			return
		// 		}
		// 		retried = true
		// 		goto retry
		// 	}
		// }(res)
		// go func( /* res *models.BetconstructData */ ) {
		// 	limiter <- true
		// }()
	}
}
