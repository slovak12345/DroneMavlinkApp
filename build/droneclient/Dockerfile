##
## Build
##
FROM golang:1.19 AS build

WORKDIR /app

COPY . ./

RUN ls

RUN go mod download

RUN go build ./cmd/droneclient/main.go

# EXPOSE 8000

ENTRYPOINT ["./cmd/droneclient/docker-entrypoint.sh"]