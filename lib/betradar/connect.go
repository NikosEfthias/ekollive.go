package betradar

import (
	"net"
	"log"
	"bufio"
)

func Connect(addr string) net.Conn {
	con, err := net.Dial("tcp", addr)
	if nil != err {
		log.Fatalln("cannot connect to betradar", err)
	}
	//con.Write([]byte(`<BookmakerStatus timestamp="0" type="login" bookmakerid="5098" key="Ea5VF7kFv" />`))
	return con
}
func GetBufferReader(con net.Conn) *bufio.Scanner {
	return bufio.NewScanner(con)
}
