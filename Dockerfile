# Building Vue SPA

FROM node:lts-alpine AS nodebuild

WORKDIR /app

COPY frontend/*.json .

RUN npm install

COPY frontend/public public
COPY frontend/src src
COPY frontend/*.js .
COPY frontend/*.ts .
COPY frontend/*.html .

RUN npm run build

# Building Go backend

FROM golang:1.23-alpine AS gobuild

WORKDIR /app

COPY --from=nodebuild /app/dist frontend/dist

COPY go.mod .
COPY main.go .
COPY internal internal
COPY logging logging

RUN CGO_ENABLED=0 GOOS=linux go build -o onyx .

# Packaging application

FROM alpine:3.21

WORKDIR /app

COPY --from=gobuild /app/onyx .

CMD [ "/app/onyx" ]
