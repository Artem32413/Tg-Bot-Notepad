package tgbot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)
type infoSave struct{
	sl [][]string
	list []string
	slEl []string
}
func Command(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	var s infoSave
	for update := range updates {
		comm := update.Message.Command()
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyMarkup = numericKeyboard
			switch comm {
			case "help":
				msg.Text = commandStr
			case "start":
				msg.Text = hello
				msg.ReplyToMessageID = update.Message.MessageID
			case "record":
				msg.Text = "Введите заметку..."
				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
					msg.Text = "Ошибка"
				}
				res, err = s.getNextValue(updates)
				if err != nil {
					log.Println(err)
					msg.Text = "Ошибка"
				} else {
					log.Printf("Значение - %s", res)
					msg.Text = "Запись успешно добавлена"
				}

			case "list":
				msg.Text = "Твои заметки: "
				if res != nil {
					for _, el := range res {
						s.slEl = append(s.slEl, el...)
					}
					for _, elSl := range s.slEl {
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
func (i infoSave) getNextValue(updates tgbotapi.UpdatesChannel) ([][]string, error) {
	for update := range updates {
		userMessage := update.Message.Text
		if !update.Message.IsCommand() {
			i.list = append(i.list, userMessage)
			i.sl = append(i.sl, i.list)
			return i.sl, nil
		}
		fmt.Println(i.sl)
	}
	return nil, nil
}

