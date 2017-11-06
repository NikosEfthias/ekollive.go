package websocketops

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"../../lib"
	"../../models/security/origin"
	ws "github.com/gorilla/websocket"
)

func StartWsServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		upg := ws.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return origin.CheckOk(r.Header.Get("origin"))
			},
		}
		con, err := upg.Upgrade(w, r, nil)
		if nil != err {
			if *lib.Testing {
				fmt.Println(err, r.Header.Get("origin"))
			}
			if !strings.Contains(err.Error(), "Origin") {
				fmt.Println(err, r.Header.Get("origin")) //If its origin error don't bother printing
			}
			return
		}
		defer con.Close()
		var old = time.Now()
		_ = old
		AddConnection(con)
		for {
			t, _, err := con.ReadMessage()
			if t == -1 || nil != err {
				DelConnection(con)
				con.Close()
				break
			}
		}
	})
	return mux
}
