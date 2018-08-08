package betradar

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"github.com/mugsoft/ekollive.go/controllers"
	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
	"github.com/mugsoft/tools/bytesize"
)

var dataChan chan *models.BetconstructData

func init() {
	dataChan = make(chan *models.BetconstructData)
}

var consuming bool
var ReadForTheFirstTime bool = false

func Parse() {
	con := Connect(*lib.ProxyURL)
	go sendping(con)
	go func() {
		var erroredNum int
		var goodNum int
		scanner := bufio.NewScanner(con)
		scanner.Split(bufio.ScanLines)
		scanner.Buffer(make([]byte, bytesize.MB*500), int(bytesize.GB))
		for scanner.Scan() {
			txt := scanner.Text()
			if len(txt) == 0 {
				// fmt.Fprintln(os.Stderr, "0 data")
				erroredNum++
				continue
			}
			var data = new(models.BetconstructData)
			// fmt.Fprintln(os.Stderr, scanner.Text())
			err := json.Unmarshal(scanner.Bytes(), data)
			if nil != err {
				// pp.Println(err)
				erroredNum++
			} else {
				goodNum++
				dataChan <- data
			}
			time.Sleep(time.Millisecond)
			// fmt.Fprintf(os.Stderr, "\033cgood => %d;bad => %d; %.2f%%", goodNum, erroredNum, (float64(erroredNum)*100.0)/float64(goodNum))
		}
		log.Println("\nbetconstruct connection was interrrupted restarting", scanner.Err())
		os.Exit(1)
	}()
	var count = 0
	for {
		count++
		select {
		case a := <-dataChan:
			if nil != a && a.Command != nil {
				controllers.Handle_data(a)
			} else {
				pp.Println(a)
			}
		case <-time.After(time.Second * 60):
			pp.Println("die motherfucker 60 seconds passed with no data")
			os.Exit(1)
		}
	}
}
func sendping(con net.Conn) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		<-ticker.C
		_, err := con.Write([]byte("ping"))
		if nil != err && os.Getenv("DEBUGGING") == "" {
			pp.Println("sock error")
			os.Exit(1)
		} else if nil != err && os.Getenv("DEBUGGING") != "" {
			ticker.Stop()
		}
	}
}
