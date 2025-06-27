# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /app/server ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 50051
ENTRYPOINT ["/app/server"]
