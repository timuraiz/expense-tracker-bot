package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	unableToSaveError = errors.New("unable to save link to Pocket")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case unableToSaveError:
		messageText = b.cfg.Errors.UnableToSave
	default:
		messageText = b.cfg.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
