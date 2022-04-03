# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as build

WORKDIR /

COPY go.mod ./
COPY go.sum ./
COPY bot.agent ./
COPY *.go ./

RUN go build -o /limbretrievalbot

##
## Deploy
##
FROM alpine:3.15.3

WORKDIR /

COPY --from=build /limbretrievalbot /etc/bin/limbretrievalbot
COPY --from=build /bot.agent /etc/bin/bot.agent

CMD [ "/etc/bin/limbretrievalbot" ]