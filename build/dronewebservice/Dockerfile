##
## Build
##
FROM golang:1.19 AS build

WORKDIR /app

COPY . ./

RUN ls

RUN go mod download

RUN go build ./cmd/dronewebservice/main.go

EXPOSE 8180/tcp

COPY ./build/dronewebservice/docker-entrypoint.sh /

RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]