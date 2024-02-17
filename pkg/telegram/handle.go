package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart      = "start"
	commandAddExpense = "add_expense"
)

type HandlerFunc func(*Bot, *tgbotapi.Message) error

var stateHandlers = map[string]HandlerFunc{
	first.GetName(): handleStartCommand,
}
var commandHandlers = map[string]HandlerFunc{
	commandStart:      handleStartCommand,
	commandAddExpense: handleAddExpenseCommand,
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return err
	}

	currState := userSession.State
	if f, exists := stateHandlers[currState.GetName()]; exists {
		return f(b, message)
	}

	f, exists := commandHandlers[message.Command()]
	if !exists {
		return handleUnknownCommand(b, message)
	}
	return f(b, message)
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	_, err := b.Bot.Send(msg)
	return err
}
