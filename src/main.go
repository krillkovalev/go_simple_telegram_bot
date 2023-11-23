package main

import (
	"os"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		if text[0] == '+' {
			text = text[1:]
		}
		if text[0] == '8' {
			text = strings.Replace(text, "8", "7", 1)
		}
		if text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, напишите номер телефона")
			bot.Send(msg)
		}

		re := regexp.MustCompile("^7\\d{10}$")
		if re.MatchString(text) == true {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://api.whatsapp.com/send?phone="+text)
			bot.Send(msg)
		} else if text != "/start" && re.MatchString(text) != true {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такой команды я не знаю, нажмите /start чтобы начать")
			bot.Send(msg)
		}

	}

}
