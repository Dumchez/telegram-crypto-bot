package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	_, err := tgbotapi.NewBotAPI("")

	if err != nil {
		log.Fatal(err)
	}

	// if err := bot.Start(); err != nil {
	// 	log.Fatal(err)
	// }

}
