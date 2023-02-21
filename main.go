//package main
//
//import (
//	"log"
//	"os"
//
//	"github.com/go-telegram-bot-api/telegram-bot-api"
//)
//
//func main() {
//	token := os.Getenv("TELEGRAM_BOT_TOKEN")
//	bot, err := tgbotapi.NewBotAPI(token)
//	if err != nil {
//		log.Panic("程序运行出错: ", err)
//	}
//
//	bot.Debug = true
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates, err := bot.GetUpdatesChan(u)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	for update := range updates {
//		if update.Message == nil {
//			continue
//		}
//
//		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//		msg.ReplyToMessageID = update.Message.MessageID
//
//		_, err := bot.Send(msg)
//		if err != nil {
//			log.Panic(err)
//		}
//	}
//}

//package main
//
//import (
//	"log"
//	"os"
//
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//)
//
//var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
//		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
//		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
//	),
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
//		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
//		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
//	),
//)
//
//func main() {
//	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates := bot.GetUpdatesChan(u)
//
//	// Loop through each update.
//	for update := range updates {
//		// Check if we've gotten a message update.
//		if update.Message != nil {
//			// Construct a new message from the given chat ID and containing
//			// the text that we received.
//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//
//			// If the message was open, add a copy of our numeric keyboard.
//			switch update.Message.Text {
//			case "open":
//				msg.ReplyMarkup = numericKeyboard
//
//			}
//
//			// Send the message.
//			if _, err = bot.Send(msg); err != nil {
//				panic(err)
//			}
//		} else if update.CallbackQuery != nil {
//			// Respond to the callback query, telling Telegram to show the user
//			// a message with the data received.
//			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
//			if _, err := bot.Request(callback); err != nil {
//				panic(err)
//			}
//
//			// And finally, send a message containing the data received.
//			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
//			if _, err := bot.Send(msg); err != nil {
//				panic(err)
//			}
//		}
//	}
//}

//package main
//
//import (
//	"log"
//	"os"
//
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//)
//
//var numericKeyboard = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("1"),
//		tgbotapi.NewKeyboardButton("2"),
//		tgbotapi.NewKeyboardButton("3"),
//	),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("4"),
//		tgbotapi.NewKeyboardButton("5"),
//		tgbotapi.NewKeyboardButton("6"),
//	),
//)
//
//func main() {
//	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates := bot.GetUpdatesChan(u)
//
//	for update := range updates {
//		if update.Message == nil { // ignore non-Message updates
//			continue
//		}
//
//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//
//		switch update.Message.Text {
//		case "open":
//			msg.ReplyMarkup = numericKeyboard
//		case "close":
//			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
//		}
//
//		if _, err := bot.Send(msg); err != nil {
//			log.Panic(err)
//		}
//	}
//}

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /code and /status."
		case "code":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm online,ready for helping."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
