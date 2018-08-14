package betradar

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

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
				fmt.Fprintln(os.Stderr, err.Error(), scanner.Text())
				erroredNum++
			} else {
				goodNum++
				dataChan <- data
			}
			time.Sleep(time.Millisecond)
			if *lib.Bar {
				fmt.Fprintf(os.Stdout, "\033cgood => %d;bad => %d; %.2f%% ==>> limiter(%v)", goodNum, erroredNum, (float64(erroredNum)*100.0)/float64(goodNum), len(controllers.Limiter))
			}
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
				fmt.Fprintln(os.Stderr, a)
			}
		case <-time.After(time.Second * 60):
			fmt.Fprintln(os.Stderr, "die  60 seconds passed with no data")
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
			fmt.Fprintln(os.Stderr, "sock error")
			os.Exit(1)
		} else if nil != err && os.Getenv("DEBUGGING") != "" {
			ticker.Stop()
		}
	}
}
