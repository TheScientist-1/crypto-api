FROM golang:1.20-alpine AS build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /CryptoAppBTCUAH

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /CryptoAppBTCUAH/btc-app .

RUN chmod +x /CryptoAppBTCUAH/btc-app

FROM scratch

COPY --from=build-stage /CryptoAppBTCUAH/btc-app /CryptoAppBTCUAH/btc-app

COPY --from=build-stage /CryptoAppBTCUAH/storage/ /CryptoAppBTCUAH/storage

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

ENTRYPOINT ["/CryptoAppBTCUAH/btc-app"]