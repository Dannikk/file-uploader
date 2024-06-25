 # syntax=docker/dockerfile:1

 ##
 ## Build the application from source
 ##

FROM golang:alpine

RUN apk update && apk add bash

WORKDIR /app

COPY ./handlers handlers
COPY ./models models
COPY ./main.go ./go.mod ./

RUN ls -al


RUN go mod tidy

RUN go build -o fileUploader

 ENTRYPOINT ["./fileUploader"]