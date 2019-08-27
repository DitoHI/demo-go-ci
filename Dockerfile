#!/bin/bash
FROM golang:alpine as build

RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

RUN mkdir -p /go/src/gitlab.com/ditohafizh/demo-go-ci

WORKDIR /go/src/gitlab.com/ditohafizh/demo-go-ci

COPY . .

RUN apk add --no-cache curl git && curl -sSL https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o /go/bin/dep && chmod +x /go/bin/dep

RUN dep ensure

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app main.go

FROM golang:alpine

COPY --from=build /go/src/gitlab.com/ditohafizh/demo-go-ci/app /usr/local/bin/app

RUN chmod +x /usr/local/bin/app

EXPOSE 9090

ENTRYPOINT ["/usr/local/bin/app"]