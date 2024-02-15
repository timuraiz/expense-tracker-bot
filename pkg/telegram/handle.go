package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart      = "start"
	commandAddExpense = "add_expense"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return handleStartCommand(b, message)
		//case commandAddExpense:
		//	return commands.HandleAddExpenseCommand(b, message, state)
		//default:
		//	return commands.HandleUnknownCommand(b, message, state)
	}
	return nil, nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	_, err := b.Bot.Send(msg)
	return err
}
