package main

import (
	"flag"
	"fmt"
	"log"
	"my-simple-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer = consumer.New()

	fmt.Println("Hello, World!")
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
