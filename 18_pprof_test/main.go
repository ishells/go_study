package main

import (
	"log"
	"net/http"
	_ "net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println(http.ListenAndServe(":6060", nil))
}
