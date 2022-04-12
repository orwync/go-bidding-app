package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orwync/go-bidding-app/auctioneer/utils"
)

func ListAPIsController(context *gin.Context) {
	listApis := []utils.ListApi{
		{
			ApiEndpoint: "/list/api",
			RequestType: "GET",
			BodyExample: "",
		},
		{
			ApiEndpoint: "/bid",
			RequestType: "POST",
			BodyExample: `{'auction_id': '123wdsasde'}`,
		},
		{
			ApiEndpoint: "/register",
			RequestType: "POST",
			BodyExample: `{'port':1000,'bidder_id':'testbidder'}`,
		},
	}
	context.JSON(http.StatusOK, &utils.ListApiResponse{
		Status: "success",
		Apis:   listApis,
		Error:  "",
	})
}
