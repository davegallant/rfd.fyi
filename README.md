# rfd-fyi

This repository provides a simple, less-distracting overlay for hot deals posted on https://forums.redflagdeals.com.

The frontend is made with Vue 3 and the backend is written in Go. The backend exists for caching purposes; to prevent excessive requests to RedFlagDeals itself.

It is hosted at [rfd.fyi](https://rfd.fyi).

## Environment Variables

Ensure that an `.env` file is filled in with the appropriate values.

Example key/values can be found in [example.env](./example.env).

*Note: that a [honeycomb](https://honeycomb.io/) API key is currently required.*

## Local Development

If running locally, ensure that the ENV vars are exported to your shell.

To get up and running locally: in one pane/tab, run:

```sh
make backend
```

In another pane/tab, run:

```sh
make frontend
```

## Docker Compose

```sh
make up
```
