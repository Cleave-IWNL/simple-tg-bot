package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	t := mustToken()

	if err != nil {
		//handle error
	}
	// token = flags.Get(token)

	// tgClient = telegram.New(token)

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
		log.Fatal("token is not specifiedS")
	}

	return *token
}
