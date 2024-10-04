FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/example.env .
COPY --from=builder /app/myapp .
RUN mv example.env .env

RUN ls -a

CMD ["./myapp"]