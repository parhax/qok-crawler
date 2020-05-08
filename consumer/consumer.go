package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/adjust/rmq"
	"qok.com/crawler/consumer/model"

	_ "github.com/go-sql-driver/mysql"
)

const (
	unackedLimit = 1000
	numConsumers = 10
	batchSize    = 1000
)

func main() {
	fmt.Printf("Consumer Started")
	redisUrl := os.Getenv("REDIS_URL")
	// connection := rmq.OpenConnection("consumer", "tcp", "localhost:6379", 1)
	connection := rmq.OpenConnection("consumer", "tcp", redisUrl, 1)
	queue := connection.OpenQueue("crawl_tasks")
	queue.StartConsuming(unackedLimit, 500*time.Millisecond)
	for i := 0; i < numConsumers; i++ {
		name := fmt.Sprintf("consumer %d", i)
		queue.AddConsumer(name, NewConsumer(i))
	}
	select {}
}

type Consumer struct {
	name   string
	count  int
	before time.Time
}

func NewConsumer(tag int) *Consumer {
	return &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
		count:  0,
		before: time.Now(),
	}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	consumer.count++
	if consumer.count%batchSize == 0 {
		duration := time.Now().Sub(consumer.before)
		consumer.before = time.Now()
		perSecond := time.Second / (duration / batchSize)
		log.Printf("%s consumed %d %s %d", consumer.name, consumer.count, delivery.Payload(), perSecond)
	}
	time.Sleep(time.Millisecond)
	if consumer.count%batchSize == 0 {
		delivery.Reject()
	} else {
		delivery.Ack()
	}
	site_url := delivery.Payload()
	crawlIt(site_url)

}

//@todo extract the crawling logic into Crawler model
func crawlIt(sUrl string) {
	resp, err := http.Get(sUrl)
	if err != nil {
		fmt.Println("An error has occured")
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read error has occured")
		} else {
			strBody := string(body)
			// re := regexp.MustCompile("<title>(.*?)</title>")
			re := regexp.MustCompile("<title*>(.*?)</title>")
			match := re.FindStringSubmatch(strBody)
			if len(match) <= 0 {
				fmt.Printf("Could not find any title for %s ", sUrl)
			} else {
				fmt.Printf("%s ", match[0])

				var crawler model.Crawler
				crawler.SetUrl(sUrl)
				crawler.SetTitle(match[0])
				crawler.StoreInDb()
			}

		}
	}
}
