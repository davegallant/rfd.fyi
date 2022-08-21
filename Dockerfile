# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

# hadolint ignore=DL3018
RUN apk --no-cache add \
  gcc \
  musl-dev

WORKDIR /app

COPY backend/ .

RUN  CGO_ENABLED=1 GOOS=linux \
    go build -o /rfd-fyi \
    # Additional flags are necessary for sqlite support
    -a -ldflags '-linkmode external -extldflags "-static"' .

EXPOSE 8080

CMD [ "/rfd-fyi" ]
