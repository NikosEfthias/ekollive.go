package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/lib/betradar"
)

func init() {
	// oddids.LoadAll()
	// filters.Init()
}
func main() {
	mux := http.NewServeMux()
	// mux.Handle("/", wso.StartWsServer()) //websocket server
	// mux.Handle("/filter/", http.StripPrefix("/filter", endpoints.Filter()))
	// mux.Handle("/command/", http.StripPrefix("/command", endpoints.Proxy()))
	if *lib.Profile {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
	}
	// var c = make(chan *models.BetconstructData)
	go betradar.Parse()
	// go wso.StartBroadcast(c)
	// go clientManager.ManageWsClients()
	fmt.Println("server listenin on port ", *lib.Port)
	log.Fatalln(http.ListenAndServe(":"+*lib.Port, mux))
}
