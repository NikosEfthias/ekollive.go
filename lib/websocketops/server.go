package websocketops

import (
	"net/http"
	ws "github.com/gorilla/websocket"
	"time"
	"fmt"
	"../../models/security/origin"
	"strings"
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
			if !strings.Contains(err.Error(), "Origin") {
			}
			fmt.Println(err, r.Header.Get("origin")) //If its origin error don't bother printing
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
				break
			}
		}
	})
	return mux
}
