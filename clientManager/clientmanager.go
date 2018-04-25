package clientManager

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type manager struct {
	write   chan []byte
	clients map[net.Conn]interface{}
}

var m = new(manager)
var l sync.Mutex

func init() {
	m.write = make(chan []byte)
	m.clients = make(map[net.Conn]interface{})
}

func ManageWsClients() {

	s, err := net.Listen("tcp", ":1111")
	if nil != err {
		panic(err)
	}
	fmt.Println("websocket manager listening on port 1111")
	go func() {
		for {
			client, err := s.Accept()
			if nil != err {
				fmt.Println("err accepting client", err)
				continue
			}
			l.Lock()
			m.clients[client] = true
			l.Unlock()
			var sigKill = make(chan interface{})
			go pingPong(client, sigKill)
			go func() {
				<-sigKill
				l.Lock()
				delete(m.clients, client)
				l.Unlock()
			}()
		}
	}()
	for {

		data := <-m.write
		for sock := range m.clients {
			_, err := sock.Write(data)
			if nil != err {
				l.Lock()
				delete(m.clients, sock)
				l.Unlock()
			}
		}
	}
}
func Broadcast(data []byte) {
	select {
	case m.write <- data:
	case <-time.After(time.Millisecond * 500):
		return
	}
}
func pingPong(c net.Conn, sigKill chan interface{}) {
	var checkAlive = make(chan interface{})
	go func() {
		for {
			var data = make([]byte, 1024)
			_, err := c.Read(data)
			if nil != err {
				c.Close()
				sigKill <- true
				break
			}
			checkAlive <- true
		}
	}()
	for {
		select {
		case <-checkAlive:
			continue
		case <-time.After(time.Millisecond * 100):
			c.Close()
			l.Lock()
			delete(m.clients, c)
			l.Unlock()
			break
		}
	}
}
