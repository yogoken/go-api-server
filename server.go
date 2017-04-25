package main

import (
	"fmt"
	"flag"
	"github.com/gin-gonic/gin"
	"golang_practice/api/v1/search"
)

func main() {
	var port = flag.Int("p", 9000, "Host IP address")
	flag.Parse()

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/search/:pattern", search.Search)
	socketaddr := fmt.Sprintf("127.0.0.1:%d", *port)
	router.Run(socketaddr)
}
