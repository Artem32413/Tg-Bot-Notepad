package tgbot

import (
	"fmt"
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
	email, password, err:= RegistrUser(updates, b.bot)
	if err != nil{
		panic(err)
	}
	fmt.Println(email, password)
	c.Command(email, password, updates, b.bot)

}
