package main

import (
	"net/http"
	"./models"
	"./lib/betradar"
	wso "./lib/websocketops"
	"./lib/store/oddids"
	"./lib/store/filters"
	"./lib"
	"flag"
)

func init() {
	lib.ApplyFlags()
	flag.Parse()
	oddids.LoadAll()
	filters.Init()
}
func main() {

	var c = make(chan models.BetradarLiveOdds)
	go betradar.Parse(c)
	go wso.StartBroadcast(c)

	http.ListenAndServe(":9090", wso.StartWsServer())
}
