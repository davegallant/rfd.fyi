FROM node:18-alpine3.16 as build-stage

WORKDIR /app

COPY package*.json ./

RUN yarn install

COPY . .

RUN yarn build

FROM caddy:2.5.2-alpine as production-stage

WORKDIR /my-site

COPY --from=build-stage /app/dist ./

COPY Caddyfile /etc/caddy/Caddyfile

