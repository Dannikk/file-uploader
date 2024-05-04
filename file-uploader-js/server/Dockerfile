 # syntax=docker/dockerfile:1

 ##
 ## Build the application from source
 ##

 FROM golang:alpine

 RUN apk update && apk add bash

 WORKDIR /app

 COPY . /app

 RUN go build -o fileUploader

 ENTRYPOINT ["./fileUploader"]