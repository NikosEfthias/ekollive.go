package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/lib/betradar"
	wso "github.com/mugsoft/ekollive.go/ws"
	"github.com/mugsoft/tools/ws"
)

func main() {
	mux := http.NewServeMux()
	if *lib.Profile {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
	}
	go betradar.Parse()
	if *lib.IsLive {
		go ws.Start_listen(wso.Opts)
	}
	log.Fatalln(http.ListenAndServe(":"+*lib.Port, mux))
}
