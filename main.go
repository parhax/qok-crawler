package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"qok.com/crawler/controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/crawl", controller.CrawlerHandler).
		Methods("POST")
	log.Fatal(http.ListenAndServe(":8686", r))
}
