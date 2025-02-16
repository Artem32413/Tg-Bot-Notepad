package tgbot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	c Commands
	b Bot
)

func Run() {
	c = &InfoSave{}
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	token := os.Getenv("TELEGRAM_APITOKEN")
	if b.bot, err = tgbotapi.NewBotAPI(token); err != nil {
		panic(err)
	}
	b.bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 60

	updates := b.bot.GetUpdatesChan(updateConfig)
	c.Command(updates, b.bot)

}
