package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adjust/rmq"
)

func CrawlerHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(req.Body)
	var arr []string
	err := json.Unmarshal(body, &arr)

	if err != nil {
		fmt.Printf("%#v", err)
	}

	connection := rmq.OpenConnection("my service", "tcp", "localhost:6379", 1)
	taskQueue := connection.OpenQueue("crawl_tasks")

	i := 0
	for _, crawl_url := range arr {
		taskQueue.Publish(crawl_url)
		fmt.Printf("%#v queued for the crawl \n ", crawl_url)
		i++
	}

	fmt.Fprintf(w, "%v sites queued for crawling", i)
	return
}
