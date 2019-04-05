package main

import (
	"github.com/heaptracetechnology/microservice-memcached/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
