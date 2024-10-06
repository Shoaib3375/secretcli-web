# Use the official Golang image as a base
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o secretcli ./cmd/secretcli.go

# Start a new stage from scratch
FROM alpine:latest
COPY --from=builder /app/secretcli .
EXPOSE 8080
CMD ["./secretcli"]
