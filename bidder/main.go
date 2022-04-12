package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/orwync/go-bidding-app/bidder/controller"
	utils "github.com/orwync/go-bidding-app/bidder/utils"
)

func main() {
	server := gin.Default()

	PORT := os.Getenv("PORT")
	BIDDING_ID := os.Getenv("BIDDER_ID")

	err := utils.RegisterToAuction(PORT, BIDDING_ID)

	if err != nil {
		fmt.Println(err.Error())
	}

	server.POST("/bidder/auction", controller.Bidding)

	// Default route error
	server.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusForbidden, gin.H{
			"error":   "invalid route",
			"message": "list endpoints at /list/api",
		})
	})

	server.Run(":" + PORT)
}
