package tgbot

import (
	"fmt"
	"log"
	// connectDB "newtgbot/pkg/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	id       string
	email    string
	password string
}

var (
	regKeybord = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/entrance"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/registration"),
		),
	)
)

func RegistrUser(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) (string, string, error) {
	customKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Войти"),
			tgbotapi.NewKeyboardButton("Зарегестрироваться"),
		),
	)
	for update := range updates {
		comm := update.Message.Command()
		if update.CallbackQuery != nil {
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			var g User
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ReplyMarkup = customKeyboard
			switch comm {
			  case "entrance", "Войти":
				msg.Text = "Введите email..."
				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
					msg.Text = "Ошибка"
				}
				email, err := g.getNextValue(updates)
				if err != nil {
					log.Println(err)
					msg.Text = "Ошибка"
				}
				if email != "" {
					msg.Text = "Введите пароль..."
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
						msg.Text = "Ошибка"
					}
					password, err := g.getNextValue2(updates)
					if err != nil {
						log.Println(err)
						msg.Text = "Ошибка"
					}
					if password != "" && email != "" {
						return email, password, nil
					}
					
				}
			default:
				msg.Text = "Команда не определена"
			}
		}

	}
	return "", "", nil
}
func (u User) getNextValue(updates tgbotapi.UpdatesChannel) (string, error) {
	for update := range updates {
		userMessage := update.Message.Text
		if !update.Message.IsCommand() {
			u.email += userMessage
			return u.email, nil
		}
		fmt.Println(u.email)
	}
	return "", fmt.Errorf("Нет данных")
}
func (u User) getNextValue2(updates tgbotapi.UpdatesChannel) (string, error) {
	for update := range updates {
		userMessage := update.Message.Text
		if !update.Message.IsCommand() {
			u.password += userMessage
			return u.password, nil
		}
		fmt.Println(u.password)
	}
	return "", fmt.Errorf("Нет данных")
}
func ConnectDB(email, password string) {

}
