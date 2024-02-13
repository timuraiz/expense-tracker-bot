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
		messageText = b.Cfg.Errors.UnableToSave
	case unableToParseExpenseError:
		messageText = b.Cfg.Errors.UnableToParse
	default:
		messageText = b.Cfg.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.Bot.Send(msg)
}
