package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/mugsoft/ekollive.go/lib"
	"github.com/mugsoft/ekollive.go/lib/betradar"
)

func init() {
}
func main() {
	mux := http.NewServeMux()
	if *lib.Profile {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
	}
	go betradar.Parse()
	log.Fatalln(http.ListenAndServe(":"+*lib.Port, mux))
}
