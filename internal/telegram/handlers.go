package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const commandStart = "start"

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleCommand(command *tgbotapi.Message) error {
	switch command.Command() {
	case commandStart:
		return b.handleStartCommand(command)
	default:
		return b.handleUnknownCommand(command)
	}
}

func (b *Bot) handleStartCommand(command *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(command.Chat.ID, "Вы ввели команду /start")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(command *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(command.Chat.ID, "Я не знаю такой команды :c")
	_, err := b.bot.Send(msg)
	return err
}
