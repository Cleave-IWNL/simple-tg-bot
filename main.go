package main

import (
	"context"
	"flag"
	"log"
	tgClient "my-simple-bot/clients/telegram"
	event_consumer "my-simple-bot/consumer/event-consumer"
	"my-simple-bot/events/telegram"
	"my-simple-bot/storage/sqlite"
)

const (
	tgBotHost = "api.telegram.org"

	storageSqlPath = "data/sqlite/storage.db"

	batchSize = 100
)

func main() {
	s, err := sqlite.New(storageSqlPath)
	if err != nil {
		log.Fatalf("can't connect to the storage: %s", err)
	}

	err = s.Init(context.TODO())

	if err != nil {
		log.Fatalf("can't innit storage: %s", err)
	}

	eventsProcessor := telegram.New(

		tgClient.New(tgBotHost, mustToken()),

		s,
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
