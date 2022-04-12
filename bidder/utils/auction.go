package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type AuctionRegister struct {
	Port      string `json:"port"`
	BiddingId string `json:"bidder_id"`
}

func RegisterToAuction(port string, biddingId string) error {
	auctionRegister, err := json.Marshal(AuctionRegister{Port: port, BiddingId: biddingId})

	resp, err := http.Post(os.Getenv("AUCTION_URL"), "application/json", bytes.NewBuffer(auctionRegister))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Printf("body: %s", string(body))

	return nil
}
