# syntax=docker/dockerfile:1.4
FROM cgr.dev/chainguard/go:latest as build

WORKDIR /src

COPY . /src

RUN CGO_ENABLED=0 go build -o backend .

FROM cgr.dev/chainguard/static@sha256:b74f66382b721769876a8dc315b89f14ea9eb914badd3a5145752d6fa2659f0c

COPY --from=build /src/backend /backend

EXPOSE 8080

CMD [ "/backend" ]
