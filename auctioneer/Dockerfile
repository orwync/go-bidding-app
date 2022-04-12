FROM golang:1.17-alpine3.14 as auctioneer-image

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Expose port
EXPOSE 4000
# Start app
CMD go run main.go