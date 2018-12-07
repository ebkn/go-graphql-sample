FROM golang:1.11.2-alpine3.8

COPY . /go/src/api

WORKDIR /go/src/api

RUN apk update --no-cache \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/...
