package main

import (
	"./controllers/endpoints"
	"./lib"
	"./lib/betradar"
	"./lib/store/filters"
	"./lib/store/oddids"
	wso "./lib/websocketops"
	"./models"
	_ "./models/language"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

func init() {
	oddids.LoadAll()
	filters.Init()
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", wso.StartWsServer()) //websocket server
	mux.Handle("/filter/", http.StripPrefix("/filter", endpoints.Filter()))
	mux.Handle("/command/", http.StripPrefix("/command", endpoints.Proxy()))
	if *lib.Profile {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
	}
	var c = make(chan models.BetradarLiveOdds)
	go betradar.Parse(c)
	go wso.StartBroadcast(c)
	fmt.Println("server listenin on port ", *lib.Port)
	log.Fatalln(http.ListenAndServe(":"+*lib.Port, mux))
}
