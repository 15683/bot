package main

import (
	"flag"
	"log"

	"github.com/15683/bot/events/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetchetm, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for acces to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
