package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/orwync/go-bidding-app/auctioneer/utils"
)

func RegisterController(context *gin.Context) {

	var bidderDetails utils.RegistrationRequestBody

	if err := context.ShouldBindJSON(&bidderDetails); err != nil {
		context.JSON(http.StatusBadRequest, utils.RegistrationResponseBody{
			Message: "unsuccessful",
			Data:    utils.RegistrationRequestBody{},
			Error:   err.Error(),
		})
		return
	}

	client := utils.ConnectRedis()

	bidders := getBidderCache(client)

	bidders = append(bidders, bidderDetails)

	fmt.Println(bidders)

	martialBidders, err := json.Marshal(bidders)

	err = client.Set("bidder", martialBidders, 0).Err()

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.RegistrationResponseBody{
			Message: "unsuccessful",
			Data:    utils.RegistrationRequestBody{},
			Error:   err.Error(),
		})
		return
	}

	context.JSON(http.StatusAccepted, utils.RegistrationResponseBody{
		Message: "success",
		Data:    bidderDetails,
		Error:   "",
	})
}

func getBidderCache(client *redis.Client) []utils.RegistrationRequestBody {
	var bidders []utils.RegistrationRequestBody
	biddersJson, err := client.Get("bidder").Result()

	json.Unmarshal([]byte(biddersJson), &bidders)

	if err != nil {
		fmt.Println(err.Error())
		return []utils.RegistrationRequestBody{}
	}

	return bidders
}
