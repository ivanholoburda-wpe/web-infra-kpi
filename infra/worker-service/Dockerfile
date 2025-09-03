FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /worker-service ./cmd/worker-service

FROM alpine:latest

COPY --from=builder /worker-service /worker-service

CMD ["/worker-service"]

