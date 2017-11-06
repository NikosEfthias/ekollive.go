package endpoints

import (
	"encoding/json"
	"net/http"

	"../../lib/endPointMethods"
	wso "../../lib/websocketops"
)

func Proxy() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		responder := json.NewEncoder(w)
		err := r.ParseForm()
		if nil != err {
			responder.Encode(Error{err})
			return
		}
		key := r.Form.Get("key")
		origin := r.Form.Get("origin")
		command := r.Form.Get("command")
		if key == "" || origin == "" || command == "" {
			responder.Encode(Error{"Missing key, origin or command"})
			return
		}
		if !endPointMethods.CheckToken(key, origin) {
			responder.Encode(Error{"invalid token or origin"})
			return
		}
		wso.Broadcast([]byte(command))
		responder.Encode(Success{true})

	})
	return mux
}
