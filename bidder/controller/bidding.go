package controller

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BiddingRequest struct {
	AuctionId string `json:"auction_id" binding:"required"`
}

type BiddingResponse struct {
	BidderId string  `json:"bidder_id"`
	Price    float64 `json:"price"`
}

type BiddingResponseBody struct {
	Message string          `json:"message"`
	Data    BiddingResponse `json:"data"`
	Error   string          `json:"error"`
}

func Bidding(context *gin.Context) {
	delay, err := strconv.Atoi(os.Getenv("DELAY"))
	if err != nil {
		delay = 100
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)

	var biddingRequest BiddingRequest

	if err := context.ShouldBindJSON(&biddingRequest); err != nil {
		context.JSON(http.StatusBadRequest, BiddingResponseBody{
			Message: "unsuccessful",
			Data:    BiddingResponse{},
			Error:   err.Error(),
		})
		return
	}

	price, err := strconv.ParseFloat(os.Getenv("PRICE"), 64)
	if err != nil {
		price = 0
	}

	context.JSON(http.StatusAccepted, BiddingResponseBody{
		Message: "success",
		Data: BiddingResponse{
			BidderId: os.Getenv("BIDDER_ID"),
			Price:    price,
		},
		Error: "",
	})

}
