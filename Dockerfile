# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-golang-webserver

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-golang-webserver /docker-golang-webserver

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-golang-webserver"]