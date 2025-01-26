package tgbot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func Run() {
	err := godotenv.Load(".env")
	if err!= nil{
		panic(err)
	}
	token := os.Getenv("TELEGRAM_APITOKEN")
	log.Println(token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err!= nil{
		panic(err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	Command(updates, bot)

}
