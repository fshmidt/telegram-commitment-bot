FROM golang:1.19-alpine3.17 AS builder

RUN go version

COPY . /github.com/fshmidt/telegram-commitment-bot/
WORKDIR /github.com/fshmidt/telegram-commitment-bot/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/fshmidt/telegram-commitment-bot/.bin/bot .
COPY --from=0 /github.com/fshmidt/telegram-commitment-bot/configs configs/
#COPY --from=0 /github.com/fshmidt/telegram-commitment-bot/bot.db .
#COPY --from=0 /github.com/fshmidt/telegram-commitment-bot/bot.db .

EXPOSE 80

CMD ["./bot"]