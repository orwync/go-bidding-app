package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AuctionRequestBody struct {
	AuctionId string `json:"auction_id"`
}

func SendAuctionRequest(bidder RegistrationRequestBody, auctionId string) (float64, string) {
	bidderJson, err := json.Marshal(AuctionRequestBody{AuctionId: auctionId})
	var biddingResult BiddingResponseBody

	startTime := time.Now()

	resp, err := http.Post(fmt.Sprintf("http://localhost:%s/bidder/auction", bidder.PortNumber), "application/json", bytes.NewBuffer(bidderJson))

	elapsedTime := time.Since(startTime).Milliseconds()

	if elapsedTime > 200 {
		return 0, ""
	}

	if err != nil {
		fmt.Println(err.Error())
		return 0, ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return 0, ""
	}

	json.Unmarshal(body, &biddingResult)

	if err != nil {
		fmt.Println(err.Error())
		return 0, ""
	}

	return biddingResult.Data.Price, bidder.BidderId
}
