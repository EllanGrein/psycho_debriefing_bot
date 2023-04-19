package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	err := b.handleUpdates(updates)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.CallbackQuery != nil {
			err := b.handleCallbackQuery(update.CallbackQuery)
			if err != nil {
				return err
			}
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message)
			if err != nil {
				return err
			}
			continue
		}
		err := b.handleMessage(update.Message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) createKeyboard() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Шкала депрессии Бека", "bek"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Тест Люшера", "lusher"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Нейрохимический определитель темперамента Фишер", "fisher"),
		),
	)

	return keyboard
}
