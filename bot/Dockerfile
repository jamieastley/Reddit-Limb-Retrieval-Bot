# syntax=docker/dockerfile:1

FROM golang:1.17.3-alpine

WORKDIR /

COPY go.mod ./
COPY go.sum ./
COPY bot.agent ./
COPY *.go ./

RUN go build -o /limbretrievalbot

CMD [ "/limbretrievalbot" ]