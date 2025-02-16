package tgbot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	
)

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("TELEGRAM_APITOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	
	Command(updates, bot)

}
