package main

import (
	"log"

	"github.com/Dumchez/telegram-crypto-bot/pkg/config"
	"github.com/Dumchez/telegram-crypto-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	tgbot := telegram.NewBot(bot)

	if err := tgbot.Start(); err != nil {
		log.Fatal(err)
	}

}
