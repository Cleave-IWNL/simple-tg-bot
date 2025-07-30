package main

import (
	"flag"
	"log"
	tgClient "my-simple-bot/clients/telegram"
	event_consumer "my-simple-bot/consumer/event-consumer"
	"my-simple-bot/events/telegram"
	"my-simple-bot/storage/files"
)

const (
	tgBotHost = "api.telegram.org"

	storagePath = "storage"

	batchSize = 100
)

// 7886272448:AAHTyVZklZ7s4RccotF4yhr6P2eTpLi6cn8

func main() {
	eventsProcessor := telegram.New(

		tgClient.New(tgBotHost, mustToken()),

		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
