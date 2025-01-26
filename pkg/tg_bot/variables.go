package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	commandStr = "Список комманд:\n/start - запуст бота\n/help - список всех комманд\n/record - начать запись, /cancel - прекратить запись\n/list - список всех заметок"
	hello      = "Привет! Это бот для работы с заметками.\nИспользуйте /help для получения списка всех доступных комманд."
	res        [][]string
	
	numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/record"),
			tgbotapi.NewKeyboardButton("/cancel"),
			tgbotapi.NewKeyboardButton("/list"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/help"),
			tgbotapi.NewKeyboardButton("/start"),
		),
	)
)

