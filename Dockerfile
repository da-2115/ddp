FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/web

RUN go mod download

RUN CGO_ENABLED=0 go build -o archery .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/web .

EXPOSE 8000

CMD ["/app/archery"]
