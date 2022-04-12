package utils

type BidRequestBody struct {
	AuctionId string `json:"auction_id" binding:"required"`
}

type BidResponse struct {
	AuctionId   string  `json:"auction_id"`
	MaxBidValue float64 `json:"maxBidValue"`
	BidderId    string  `json:"bidder_id"`
}

type BidResponseBody struct {
	Message string      `json:"message"`
	Data    BidResponse `json:"data"`
	Error   string      `json:"error"`
}

type ListApi struct {
	ApiEndpoint string `json:"apiEndpoint"`
	RequestType string `json:"requestType"`
	BodyExample string `json:"bodyExample"`
}

type ListApiResponse struct {
	Status string    `json:"status"`
	Apis   []ListApi `json:"apis"`
	Error  string    `json:"error"`
}

type RegistrationRequestBody struct {
	BidderId   string `json:"bidder_id" binding:"required"`
	PortNumber string `json:"port" binding:"required"`
}

type RegistrationResponseBody struct {
	Message string                  `json:"message"`
	Data    RegistrationRequestBody `json:"data"`
	Error   string                  `json:"error"`
}

type BidBody struct {
	BidValue float64 `json:"bid"`
	BidderId string  `json:"bidder_id"`
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
