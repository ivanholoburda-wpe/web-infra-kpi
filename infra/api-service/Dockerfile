FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-service ./cmd/api-service

FROM alpine:latest

COPY --from=builder /api-service /api-service

EXPOSE 8080

CMD ["/api-service"]

