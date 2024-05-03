FROM golang:1.22.2-alpine AS builder

WORKDIR /github.com/go-jedi/go-telegram-clean-architecture/app/
COPY . /github.com/go-jedi/go-telegram-clean-architecture/app/

RUN go mod download
RUN go build -o .bin/telegram_bot cmd/telegram_bot/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/go-telegram-clean-architecture/app/.bin/telegram_bot .
COPY .env /root/

CMD ["./telegram_bot"]