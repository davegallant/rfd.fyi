FROM node:20.1.0-alpine3.16 as builder

WORKDIR /app

COPY package*.json ./

RUN yarn install

COPY . .

RUN yarn build

FROM caddy:2.6.4-alpine as runtime

WORKDIR /my-site

COPY --from=builder /app/dist ./

COPY Caddyfile /etc/caddy/Caddyfile

