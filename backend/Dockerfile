# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as go-build

# hadolint ignore=DL3018
RUN apk --no-cache add \
  gcc \
  musl-dev

COPY . /backend

WORKDIR /backend

RUN  CGO_ENABLED=1 GOOS=linux \
    go build -o server \
    # Additional flags are necessary for sqlite support
    -a -ldflags '-linkmode external -extldflags "-static"' .

EXPOSE 8080

CMD [ "/backend/server" ]
