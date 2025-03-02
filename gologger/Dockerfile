# Use the official Go image as the base image
FROM golang:1.23.6-alpine  AS builder
WORKDIR /app
COPY main.go .
RUN go mod init gologger && go build -o gologger

FROM alpine:latest
RUN apk add --no-cache netcat-openbsd
WORKDIR /app
COPY --from=builder /app/gologger .
EXPOSE 3001
CMD ["./gologger"]