# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o schedule_gateway ./main.go

# stage 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/schedule_gateway .
RUN chmod +x /app/schedule_gateway
ENTRYPOINT ["./schedule_gateway"]
