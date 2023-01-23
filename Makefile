PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-commitment-bot:v0.1 .
start-container:
	docker run --name telegram-commitment-bot -p 80:80 --env-file .env telegram-commitment-bot:v0.1