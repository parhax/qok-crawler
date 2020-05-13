package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adjust/rmq"
	"qok.com/crawler/http/config"
	"qok.com/crawler/http/logwrapper"
)

func CrawlerHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		logwrapper.Load().Fatal("ioutil read error in crawlerHandler")
	}
	var arr []string
	err := json.Unmarshal(body, &arr)

	if err != nil {
		logwrapper.Load().Fatalf("%#v", err)
	}
	redisUrl := config.Load().Redis_url
	connection := rmq.OpenConnection("my service", "tcp", redisUrl, 1)
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
