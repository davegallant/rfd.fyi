# rfd.fyi

This repository provides a simple, less-distracting overlay for hot deals posted on https://forums.redflagdeals.com.

The frontend is made with Vue 3 and the backend is written in Go. The backend exists for caching purposes; to prevent excessive requests to RedFlagDeals itself.

It is hosted at [rfd.fyi](https://rfd.fyi).

## Local Development

To get up and running locally: in one pane/tab, run:

```sh
make backend
```

In another pane/tab, run:

```sh
make frontend
```

## Docker Compose

To build containers from source:

```sh
make dev
```

To run the latest published images:

```sh
make prod
```
