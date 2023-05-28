FROM golang:1.20-alpine AS build-stage

WORKDIR /Crypto_API

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o /Crypto_API/main .

FROM alpine:3.18

COPY --from=build-stage /Crypto_API/main main

EXPOSE 8080

ENTRYPOINT ["main"]