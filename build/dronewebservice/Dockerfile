##
## Build
##
FROM golang:1.19 AS build

WORKDIR /app

COPY . ./

RUN ls

RUN go mod download

RUN go build ./cmd/dronewebservice/main.go

# EXPOSE 8000

ENTRYPOINT ["./cmd/dronewebservice/docker-entrypoint.sh"]