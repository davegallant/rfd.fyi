# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as go-build

COPY . /backend

WORKDIR /backend

RUN CGO_ENABLED=0 go build -o server

EXPOSE 8080

CMD [ "/backend/server" ]
