# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as build

WORKDIR /src

RUN mkdir -p api
RUN mkdir -p repository

COPY api/ ./api/
COPY repository/ ./repository/
COPY go.work ./
COPY go.work.sum ./

WORKDIR api

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o build

##
## Deploy
##
FROM alpine:3.15.3

WORKDIR /

COPY --from=build /src/api/build /etc/bin/api

RUN ls /etc/bin/api

CMD [ "/etc/bin/api" ]