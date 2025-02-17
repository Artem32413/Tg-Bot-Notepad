package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Commands interface {
	Command(email, password string, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI)
}
type InfoSave struct {
	sl   [][]string
	list []string
	slEl []string
}

type Bot struct {
	bot *tgbotapi.BotAPI
}


var (
	commandStr = "Список комманд:\n/start - запуст бота\n/help - список всех комманд\n/record - Введите заметку...\n/list - список всех заметок\n"
	hello      = "Привет! Это бот для работы с заметками.\nИспользуйте /help для получения списка всех доступных комманд."
	err        error
	res        [][]string
	slEl       []string

	numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/record"),
			tgbotapi.NewKeyboardButton("/list"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/help"),
			tgbotapi.NewKeyboardButton("/profil"),
			tgbotapi.NewKeyboardButton("/start"),
		),
	)
	
)
