package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart       = "start"
	callbackDataLusher = "lusher"
	callbackDataBek    = "bek"
	callbackDataFisher = "fisher"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, введите команду /start и выберите психологический тест")
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
	msg := tgbotapi.NewMessage(command.Chat.ID, "Привет! Выбери тест, который хочешь пройти:")

	keyboard := b.createKeyboard()
	msg.ReplyMarkup = keyboard

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(command *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(command.Chat.ID, "Я не знаю такой команды :c")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCallbackQuery(callback *tgbotapi.CallbackQuery) error {
	switch callback.Data {
	case callbackDataBek:
		return b.handleBekCallback(callback.Message.Chat.ID)
	case callbackDataLusher:
		return b.handleLusherCallback(callback.Message.Chat.ID)
	case callbackDataFisher:
		return b.handleFisherCallback(callback.Message.Chat.ID)
	default:
		return b.handleUnknownCallback(callback.Message.Chat.ID)
	}
}

func (b *Bot) handleBekCallback(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, "Тест \"Шкала депрессии Бека\"")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleLusherCallback(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, "Тест \"Тест Люшера\"")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleFisherCallback(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, "Тест \"Нейрохимический определитель темперамента Фишер\"")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCallback(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, "Я не знаю такого опроса :(")
	_, err := b.bot.Send(msg)
	return err
}
