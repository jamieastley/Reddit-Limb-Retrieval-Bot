# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as build

WORKDIR /

COPY go.mod ./
COPY go.sum ./
COPY graph ./graph
COPY gqlgen.yml ./
COPY *.go ./

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /api

##
## Deploy
##
FROM alpine:3.15.3

WORKDIR /

COPY --from=build /api /etc/bin/api

CMD [ "/etc/bin/api" ]