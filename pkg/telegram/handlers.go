package telegram

import (
	"fmt"

	"github.com/Dumchez/telegram-crypto-bot/pkg/crypto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	if message.Command() == "help" {
		msg.Text = commands
		b.bot.Send(msg)
		return nil
	}

	price, err := crypto.GetCryptoPrice(message.Command())
	if err != nil {
		return err
	}
	if price == 0 {
		msg.Text = "There are no such coins in binance database. \nIf you think you have entered it correctly, check the spelling! All coins should be in upper case: BTC"
		b.bot.Send(msg)
		return nil
	}

	msg.Text = fmt.Sprintf("The price of %s is %0.2f$", message.Command(), price)
	b.bot.Send(msg)
	return nil
}
