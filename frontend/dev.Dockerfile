
FROM node:latest

WORKDIR /usr/src/govulcantv/frontend
COPY package.json yarn.lock ./
RUN yarn install --network-timeout 10000000
COPY src src
COPY public public
CMD yarn start
