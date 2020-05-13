package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"qok.com/crawler/http/controller"
)

func Run(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/crawl", controller.CrawlerHandler).
		Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
