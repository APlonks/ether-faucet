package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	fmt.Println(c)
	fmt.Println("TEST")
}

func main() {

	// Create client and all

	router := gin.Default()

	router.GET("/someGet", getting)

	router.POST("/SendEthersToSpecificAddress", func(ctx *gin.Context) {
		// faucet.SendEthersToSpecificAddress()
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
