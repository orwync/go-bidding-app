package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/orwync/go-bidding-app/auctioneer/controllers"
)

func main() {
	server := gin.Default()

	server.GET("/list/api", controllers.ListAPIsController)

	server.POST("/bid", controllers.BidController)

	server.POST("/registration", controllers.RegisterController)

	// Default route error
	server.NoRoute(controllers.DefaultController)

	server.Run(":4000")
}
