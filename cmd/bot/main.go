package main

import (
	"github.com/EllanGrein/psycho_debriefing_bot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
