FROM golang:buster

RUN mkdir /deploy
WORKDIR /deploy

COPY . /deploy

RUN go mod download
