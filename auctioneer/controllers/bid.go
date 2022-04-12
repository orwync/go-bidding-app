package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orwync/go-bidding-app/auctioneer/utils"
)

func BidController(context *gin.Context) {

	registerBody := utils.BidRequestBody{}
	maxBidder := utils.BidBody{}

	if err := context.ShouldBindJSON(&registerBody); err != nil {
		context.JSON(http.StatusBadRequest, utils.BidResponseBody{
			Message: "unsuccessful",
			Data:    utils.BidResponse{},
			Error:   err.Error(),
		})
		return
	}

	client := utils.ConnectRedis()
	bidders := utils.GetBidderCache(client)

	client.FlushAll()

	for _, bidder := range bidders {
		bids, bidderId := utils.SendAuctionRequest(bidder, registerBody.AuctionId)
		if maxBidder.BidValue < bids {
			maxBidder.BidValue = bids
			maxBidder.BidderId = bidderId

		}
	}

	if maxBidder.BidValue == 0 && maxBidder.BidderId == "" {
		context.JSON(http.StatusPartialContent, utils.BidResponseBody{
			Message: "no bids",
			Data: utils.BidResponse{
				AuctionId:   registerBody.AuctionId,
				MaxBidValue: 0,
				BidderId:    "",
			},
			Error: "",
		})

		return
	}

	context.JSON(http.StatusAccepted, utils.BidResponseBody{
		Message: "success",
		Data: utils.BidResponse{
			AuctionId:   registerBody.AuctionId,
			MaxBidValue: maxBidder.BidValue,
			BidderId:    maxBidder.BidderId,
		},
		Error: "",
	})
}
