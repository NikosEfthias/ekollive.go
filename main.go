package main

import (
	"net/http"
	_ "net/http/pprof"

	"./lib"
	"./lib/betradar"
	"./lib/store/filters"
	"./lib/store/oddids"
	wso "./lib/websocketops"
	"./models"
	"log"
	"fmt"
)

func init() {
	oddids.LoadAll()
	filters.Init()
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", wso.StartWsServer()) //websocket server
	var c = make(chan models.BetradarLiveOdds)
	go betradar.Parse(c)
	go wso.StartBroadcast(c)
	fmt.Println("server listenin on port ", *lib.Port)
	log.Fatalln(http.ListenAndServe(":" + *lib.Port, mux))
}
