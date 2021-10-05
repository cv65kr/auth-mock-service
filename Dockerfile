FROM golang:1.17-alpine as builder

RUN mkdir -p /app/

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -a -o bin/main cmd/main/*

FROM alpine:3.14

WORKDIR /home/app

COPY --from=builder /app/bin/main .

EXPOSE 8000

CMD ["./main"]