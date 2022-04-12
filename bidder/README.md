# Go Bidder Service

Go service to bid for auction

## Usage

To change the parameters change in docker-compose.yaml file as shown below for any of the bidders.

```
- PORT=1001
- BIDDER_ID=id1
- PRICE=1000
- DELAY=100
- AUCTION_URL=http://localhost:4000/registration
```

Open the bidder directory in the terminal and run : 
```docker-compose up```

Make sure no applications are running on the port 4000 (app) and 6379 (redis docker image)