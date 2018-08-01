package betradar

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/mugsoft/tools"
)

func Connect(addr string) net.Conn {
	con, err := net.Dial("tcp", addr)
	if nil != err {
		log.Fatalln("cannot connect to betradar", err)
	}
	Login(con)
	return con
}

type Command struct {
	Command string                   `json:"Command"`
	Params  []map[string]interface{} `json:"Params"`
	Objects []map[string]interface{} `json:"Objects,omitempty"`
}

func (l *Command) Send(sock net.Conn) error {
	d, err := json.Marshal(l)
	if nil != err {
		return err
	}
	ln := tools.Int2LE(uint(len(d)))
	_, err = sock.Write(ln[:])
	if nil != err {
		return err
	}
	_, err = sock.Write(d)
	return err
}

func LoginWithValues(uname string, pass string) *Command {
	return &Command{
		Command: "Login",
		Params:  []map[string]interface{}{{"UserName": uname, "Password": pass}},
	}
}
func Login(sock net.Conn) {
	LoginWithValues("demouser6", "4AS5V_a;wW").Send(sock)
	var cmds = []*Command{
		{
			Command: "GetSports",
		},
		{
			Command: "GetRegions",
		},
		{
			Command: "GetCompetitions",
		},
		{
			Command: "GetMarketTypes",
			Params:  []map[string]interface{}{{"SportId": "1"}},
		},
	}
	for _, cmd := range cmds {
		// time.Sleep(time.Second)
		cmd.Send(sock)
		go func() {
			for {
				(&Command{Command: "HeartBeat"}).Send(sock)
				time.Sleep(time.Second * 5)
			}
		}()
	}
}
