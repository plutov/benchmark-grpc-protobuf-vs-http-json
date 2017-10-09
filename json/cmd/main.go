package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/plutov/benchmark-grpc-vs-json/json"
)

func main() {
	json.Start()

	log.Println(http.ListenAndServe("localhost:6060", nil))
}
