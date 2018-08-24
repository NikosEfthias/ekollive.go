package betradar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"github.com/mugsoft/ekollive.go/controllers"
	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/models"
	"github.com/mugsoft/tools"
	"github.com/mugsoft/tools/bytesize"
)

var dataChan chan *models.Command

func init() {
	dataChan = make(chan *models.Command)
}

var consuming bool
var ReadForTheFirstTime bool = false
var buffering bool

func Parse() {
	con := Connect(*lib.Lang_api_url)
	go Login(con)
	go sendping(con)
	go func() {
		// var erroredNum int
		// var goodNum int
		const bufsize = bytesize.KB * 500
		var buf = make([]byte, bufsize)
		var firstTime = true
		var fullData = []byte{}
		for {
			if !firstTime {
			} else {
				firstTime = false
			}
			var length int
			var remaining int
			var meta = make([]byte, 4)
			fullData = fullData[:0]
			n, err := con.Read(meta)
			if nil != err {
				fmt.Println("read error")
				fmt.Println(err)
				break
			} else if n < 4 {
				fmt.Println("Erroorrr they sent less bytes ")
				continue
			}
			// scanChan <- meta
			length = int(tools.LE2Int(meta))
			remaining = length
			for remaining > 0 {
				buffering = true
				if uint64(length) > bytesize.MB*30 {
					pp.Println(">>", remaining, len(buf))
				}
				if remaining < len(buf) {
					buf = buf[:remaining]
				} else if remaining > int(bufsize) && len(buf) < int(bufsize) {
					buf = make([]byte, bufsize)
				}
				n, _ = con.Read(buf)
				remaining -= n
				// scanChan <- buf[:n]
				fullData = append(fullData, buf[:n]...)
			}
			var data = new(models.Command)

			err = json.Unmarshal(fullData, data)
			if nil != err {
				log__file("errored_json.json", os.Stderr, string(fullData))
				fmt.Fprintln(os.Stderr, "err", err)
				continue
			}
			dataChan <- data
			buffering = false
		}
		log.Println("\nbetconstruct connection was interrrupted restarting")
		os.Exit(1)
	}()
	for {
		select {
		case a := <-dataChan:
			controllers.Add_me_if_you_can(a)
		case <-time.After(time.Second * 60):
			if !buffering {
				fmt.Fprintln(os.Stderr, "die  60 seconds passed with no data")
				os.Exit(1)
			}
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
func log__file(fname string, fallback io.Writer, data ...interface{}) {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if nil != err {
		fmt.Fprintln(fallback, data...)
		return
	}
	fmt.Fprintln(file, data...)
	file.Close()

}

func LoginWithValues(uname string, pass string) *models.Command {
	return &models.Command{
		Command: "Login",
		Params:  []map[string]interface{}{{"UserName": uname, "Password": pass}},
	}
}

func Login(sock net.Conn) {
	LoginWithValues(*lib.Key, *lib.Pass).Send(sock)
	go send__cmds(sock)
}
func send__cmds(sock net.Conn) {
	f, err := os.Open("lang-cmds.json")
	if nil != err {
		fmt.Fprintln(os.Stderr, "error opening lang file lang-cmds.json", err.Error())
		os.Exit(1)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var jsnstr = &struct {
		Commands  []*models.Command `json:"commands"`
		Endpoints []string          `json:"endpoints"`
	}{}
	err = dec.Decode(jsnstr)
	if nil != err {
		fmt.Fprintln(os.Stderr, "invalid json file", err.Error())
		os.Exit(1)
	}
	for _, cmd := range jsnstr.Commands {
		time.Sleep(time.Second)
		fmt.Println("sending cmd:", cmd.Command)
		cmd.Timeout = 0
		cmd.Send(sock)
		timer(cmd.Timeout)
	}
}
func timer(rmn int) {
	for rmn > 0 {
		fmt.Print("\r      \r")
		fmt.Print(rmn)
		time.Sleep(time.Second)
		rmn--
	}
}
