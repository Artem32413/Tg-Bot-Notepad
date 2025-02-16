package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
type List struct{
	Id int `json:"id"`
	Text string `json:"text"`
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
			tgbotapi.NewKeyboardButton("/start"),
		),
	)
)
