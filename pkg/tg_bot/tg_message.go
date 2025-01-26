package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Message(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
