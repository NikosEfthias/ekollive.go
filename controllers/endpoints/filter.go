package endpoints

import (
	"../../lib/endPointMethods"
	"../../lib/store/filters"
	"encoding/json"
	"net/http"
)

func Filter() *http.ServeMux {
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
		filter := r.Form.Get("filter")
		if key == "" || origin == "" || filter == "" {
			responder.Encode(Error{"Missing key, origin or filter"})
			return
		}
		if !endPointMethods.CheckToken(key, origin) {
			responder.Encode(Error{"invalid token or origin"})
			return
		}
		if err := filters.Add(filter); nil != err {
			responder.Encode(Error{err.Error()})
			return
		}
		responder.Encode(Success{true})
	})
	mux.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		filters.LoadAll()
		responder := json.NewEncoder(w)
		err := r.ParseForm()
		if nil != err {
			responder.Encode(Error{err})
			return
		}
		key := r.Form.Get("key")
		origin := r.Form.Get("origin")
		if key == "" || origin == "" {
			responder.Encode(Error{"Missing key or origin"})
			return
		}
		if !endPointMethods.CheckToken(key, origin) {
			responder.Encode(Error{"invalid token or origin"})
			return
		}
		responder.Encode(Success{true})
	})
	return mux
}
