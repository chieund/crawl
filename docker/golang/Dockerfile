FROM golang:1.18

RUN apt update && apt -y upgrade

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

EXPOSE 9000