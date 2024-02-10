package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"strconv"
	"strings"
	"time"
)

const (
	commandStart      = "start"
	commandAddExpense = "add_expense"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandAddExpense:
		return b.handleAddExpenseCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.cfg.Responses.Start)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.cfg.Responses.UnknownCommand)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddExpenseCommand(message *tgbotapi.Message) error {
	content := strings.Fields(message.Text)[1:]
	var amountText, categoryText string
	if len(content) == 2 {
		amountText, categoryText = content[0], content[1]
	} else {
		return unableToParseExpenseError
	}
	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		return unableToParseExpenseError
	}

	expenseDetail := storage.ExpenseDetail{
		UserID:      message.Chat.ID,
		Amount:      amount,
		Category:    categoryText,
		Date:        time.Now(),
		Description: "",
	}
	_, err = b.db.AddExpense(expenseDetail)
	if err != nil {
		return unableToSaveExpenseError
	}
	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.cfg.Responses.ExpenseSaved))

	return err
}
