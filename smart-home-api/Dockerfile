# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN go get

RUN go build -o /smart-home-api

EXPOSE 8080

CMD [ "/smart-home-api" ]