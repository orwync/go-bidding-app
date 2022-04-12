package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	return client
}

func GetBidderCache(client *redis.Client) []RegistrationRequestBody {
	var bidders []RegistrationRequestBody
	biddersJson, err := client.Get("bidder").Result()

	json.Unmarshal([]byte(biddersJson), &bidders)

	if err != nil {
		fmt.Println(err.Error())
		return []RegistrationRequestBody{}
	}

	return bidders
}
