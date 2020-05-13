package main

import (
	"qok.com/crawler/http/config"
	"qok.com/crawler/http/router"
)

func main() {
	router.Run(config.Load().Http_port)
}
