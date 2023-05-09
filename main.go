package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6274923928:AAFK-hEM6HA2LrOTyJFxNnATnK6hS_-40VE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Создаем кнопки для подменю
	btn1 := tgbotapi.NewKeyboardButton("добавить валюту")
	btn2 := tgbotapi.NewKeyboardButton("удалить валюту")
	btn3 := tgbotapi.NewKeyboardButton("баланс онлайн")
	btn4 := tgbotapi.NewKeyboardButton("баланс не онлайн")
	btn5 := tgbotapi.NewKeyboardButton("Сводка сегодня")
	btn6 := tgbotapi.NewKeyboardButton("Удалить сводку крипты")
	btn7 := tgbotapi.NewKeyboardButton("Анализ финансов сегодня")
	btn8 := tgbotapi.NewKeyboardButton("Анализ финансов завтра")

	// Создаем массив кнопок для подменю
	row := []tgbotapi.KeyboardButton{btn1, btn2}
	lop := []tgbotapi.KeyboardButton{btn3, btn4}
	sop := []tgbotapi.KeyboardButton{btn5, btn6}
	dop := []tgbotapi.KeyboardButton{btn7, btn8}

	// Создаем клавиатуру для подменю
	submenu := tgbotapi.NewReplyKeyboard(row)
	submenu2 := tgbotapi.NewReplyKeyboard(lop)
	submenu3 := tgbotapi.NewReplyKeyboard(sop)
	submenu4 := tgbotapi.NewReplyKeyboard(dop)

	// Создаем кнопку для основного меню
	mainBtn := tgbotapi.NewKeyboardButton("Кошелек")
	mainBtn2 := tgbotapi.NewKeyboardButton("Электронный кошелек")
	mainBtn3 := tgbotapi.NewKeyboardButton("Сводка крипты")
	mainBtn4 := tgbotapi.NewKeyboardButton("Анализ рынка")

	// Создаем массив кнопок для основного меню
	mainRow := []tgbotapi.KeyboardButton{mainBtn}
	mainRow2 := []tgbotapi.KeyboardButton{mainBtn2}
	mainRow3 := []tgbotapi.KeyboardButton{mainBtn3}
	mainRow4 := []tgbotapi.KeyboardButton{mainBtn4}

	// Создаем клавиатуру для основного меню
	mainMenu := tgbotapi.NewReplyKeyboard(mainRow, mainRow2, mainRow3, mainRow4)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {

		case "Кошелек":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт меню")
			msg.ReplyMarkup = submenu
			bot.Send(msg)
		case "добавить валюту":

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Валюта добавлена! ")
			msg.ReplyMarkup = submenu
			bot.Send(msg)
		case "удалить валюту":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Валюта удалена")
			msg.ReplyMarkup = submenu
			msg.ReplyMarkup = mainMenu
			bot.Send(msg)

		// Работает адекватно

		case "Электронный кошелек":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт меню")
			msg.ReplyMarkup = submenu2
			bot.Send(msg)
		case "баланс онлайн":

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Баланс онлайн 0.00")
			msg.ReplyMarkup = submenu2
			bot.Send(msg)
		case "баланс не онлайн":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "мой хозяин не предпологал,что сюда писать")
			msg.ReplyMarkup = submenu2
			msg.ReplyMarkup = mainMenu
			bot.Send(msg)

			// кнопка 3
		case "Сводка крипты":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт меню")
			msg.ReplyMarkup = submenu3
			bot.Send(msg)

		case "Сводка сегодня":

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мой хозяин не добавил эту услугу:с")
			msg.ReplyMarkup = submenu3
			bot.Send(msg)

		case "Удалить сводку крипты":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "мой хозяин не предпологал,что сюда писать")
			msg.ReplyMarkup = submenu3
			msg.ReplyMarkup = mainMenu
			bot.Send(msg)
			// кнопка 4
		case "Анализ рынка":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите пункт меню")
			msg.ReplyMarkup = submenu4
			bot.Send(msg)

		case "Анализ финансов сегодня":

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мой хозяин не добавил эту услугу:с")
			msg.ReplyMarkup = submenu4
			bot.Send(msg)

		case "Анализ финансов завтра":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "мой хозяин не предпологал,что сюда писать")
			msg.ReplyMarkup = submenu4
			msg.ReplyMarkup = mainMenu
			bot.Send(msg)

		default:
			// Отправляем сообщение с основным меню
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, ребята из Вконтакте:")
			msg.ReplyMarkup = mainMenu
			bot.Send(msg)

		}

	}
}
