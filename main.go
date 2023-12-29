package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		botToken = "OR_FILL_TOKEN_HERE"
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Kiểm tra nếu đây là tin nhắn từ một nhóm (Group) hoặc siêu nhóm (Supergroup)
		if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
			// Lấy thông tin về người gửi tin nhắn
			sender := update.Message.From

			// Hiển thị nội dung tin nhắn và thông tin về người gửi
			index := strings.Index(update.Message.Text, " ")
			log.Printf("[%s] %s", sender.UserName, update.Message.Text[index+1:])
		}
	}
}
