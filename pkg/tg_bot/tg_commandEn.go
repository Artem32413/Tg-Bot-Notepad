package tgbot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func Command(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	var err error
	var slEl []string
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyMarkup = numericKeyboard 
			switch update.Message.Command() {
			case "help":
				msg.Text = commandStr
			case "start":
				msg.Text = hello
				msg.ReplyToMessageID = update.Message.MessageID
			case "record":
				res, err = getNextValue(updates)
				if err != nil {
					log.Println(err)
					msg.Text = "Ошибка"
				}
				log.Printf("Значение - %s", res)
				msg.Text = "Запись успешно добавлена"
			case "list":
				if res != nil {
					for _, el := range res {
						slEl = append(slEl, el...)
					}
					for _, elSl:= range slEl{
						msg.Text += elSl
					}
				} else {
					msg.Text = "Список заметок пуст"
				}
			default:
				msg.Text = "Команда не определена"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
				msg.Text = "Ошибка"
			}
			continue
		}
	}
}
func getNextValue(updates tgbotapi.UpdatesChannel) ([][]string, error) {
	var sl [][]string
	var list []string
	for update := range updates {

		userMessage := update.Message.Text
		if !update.Message.IsCommand() {
			list = append(list, userMessage)
			sl = append(sl, list)
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
