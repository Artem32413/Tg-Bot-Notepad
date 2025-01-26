package tgbot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	commandStr = "/start - запуст бота\n/help - список всех комманд\n/record - записаться, /cancel - прекратить запись\n"
)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
)

func Command(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "Список комманд:\n"
				msg.Text = commandStr
			case "start":
				msg.Text = "Привет! Это бот для работы с заметками.\nИспользуйте /help для получения списка всех доступных комманд."
				msg.ReplyToMessageID = update.Message.MessageID
			case "record":
				res, err := getNextValue(updates)
				if err != nil {
					log.Println(err)
					msg.Text = "Ошибка"
				}
				log.Printf("Значение - %s", res)
				msg.Text = "Запись успешно добавлена"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
				msg.Text = "Ошибка"
			}
			continue
		}

		Message(update, bot)
	}
}
func getNextValue(updates tgbotapi.UpdatesChannel) ([]string, error) {
	var sl []string
	for update := range updates {
		userMessage := update.Message.Text
		if !update.Message.IsCommand() {
			sl = append(sl, userMessage)

		} else {
			switch update.Message.Command() {
			case "cancel":
				return sl, nil
			}
		}
		fmt.Println(sl)
	}

	return nil, nil
}
