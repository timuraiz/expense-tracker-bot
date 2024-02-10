package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	unableToSaveExpenseError  = errors.New("unable to save expense")
	unableToParseExpenseError = errors.New("unable to parse expense")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case unableToSaveExpenseError:
		messageText = b.cfg.Errors.UnableToSave
	case unableToParseExpenseError:
		messageText = b.cfg.Errors.UnableToParse
	default:
		messageText = b.cfg.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
