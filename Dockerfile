FROM node:20.2.0-alpine3.16 AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM caddy:2.8.4-alpine AS runtime

WORKDIR /my-site

COPY --from=builder /app/dist ./

COPY Caddyfile /etc/caddy/Caddyfile

