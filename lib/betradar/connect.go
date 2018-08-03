package betradar

import (
	"log"
	"net"
)

func Connect(addr string) net.Conn {
	con, err := net.Dial("tcp", addr)
	if nil != err {
		log.Fatalln("cannot connect to betradar", err)
	}
	return con
}
