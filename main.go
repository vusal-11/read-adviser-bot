package main

import (
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	event_consumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

// 7247711279:AAFIIpniUGIqmGUFZqan9YdrzW06u9DsHyw

func main() {

	//TODO

	// token = flags.Get(token)

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Print("services started")
	// fetcher =  fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher,processor)

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("services is stopped", err)
	}
}

func mustToken() string {
	//bot -tg-bot-token  'my token'
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
