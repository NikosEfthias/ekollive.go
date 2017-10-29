package main

import (
	"./models"
	"fmt"
	"net/http"
	"net/http/pprof"
	"log"
	wso "./lib/websocketops"
	"./lib/betradar"
	"./controllers/endpoints"
	"./lib"
	"./lib/store/filters"
	"./lib/store/oddids"
)

func init() {
	oddids.LoadAll()
	filters.Init()
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", wso.StartWsServer()) //websocket server
	mux.Handle("/filter/", http.StripPrefix("/filter", endpoints.Filter()))
	if *lib.Profile {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
	}
	var c = make(chan models.BetradarLiveOdds)
	go betradar.Parse(c)
	go wso.StartBroadcast(c)
	fmt.Println("server listenin on port ", *lib.Port)
	log.Fatalln(http.ListenAndServe(":" + *lib.Port, mux))
}
