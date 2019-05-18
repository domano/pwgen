# Build pwgen as a stand alone binary
FROM golang:latest as builder
COPY . /pwgen
WORKDIR /pwgen
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o pwgen cmd/pwgen/main.go

# Run pwgen as non-root user in small container
FROM alpine:latest
RUN apk add --no-cache tzdata
RUN addgroup -S pwgen && adduser -S -G pwgen pwgen
USER pwgen
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /pwgen/pwgen .
EXPOSE 8443/tcp
CMD ["./pwgen"]