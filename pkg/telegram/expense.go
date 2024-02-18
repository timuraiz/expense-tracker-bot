package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"github.com/timuraiz/expense-tracker-bot/pkg/telegram/session"
	"strconv"
	"time"
)

var (
	askMoneyCost = session.NewState("askMoneyCost")
	saveExpense  = session.NewState("saveInDb")
)

func handleAddExpenseCommand(b *Bot, message *tgbotapi.Message) error {
	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return unableToReachUserSessionError
	}
	switch userSession.State.GetName() {
	case askMoneyCost.GetName():
		return handleAskMoneyCost(b, message)
	case saveExpense.GetName():
		return handleSaveExpense(b, message)
	default:
		return handleAskCategory(b, message)
	}

}

func handleAskCategory(b *Bot, message *tgbotapi.Message) error {
	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return unableToReachUserSessionError
	}
	_, err = b.Bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.ExpenseCategory))
	if err != nil {
		return err
	}
	userSession.SetState(&askMoneyCost)
	return nil
}

func handleAskMoneyCost(b *Bot, message *tgbotapi.Message) error {
	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return unableToReachUserSessionError
	}

	userSession.Data["category"] = message.Text

	_, err = b.Bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.ExpenseMoneyCost))
	if err != nil {
		return err
	}
	userSession.SetState(&saveExpense)
	return nil
}

func handleSaveExpense(b *Bot, message *tgbotapi.Message) error {
	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(message.Text, 64)
	if err != nil {
		return unableToParseExpenseError
	}
	category := userSession.Data["category"].(string)
	expenseDetail := storage.ExpenseDetail{
		UserID:      message.Chat.ID,
		Amount:      amount,
		Category:    category,
		Date:        time.Now(),
		Description: "",
	}
	_, err = b.Db.AddExpense(expenseDetail)
	if err != nil {
		return unableToSaveExpenseError
	}
	_, err = b.Bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.ExpenseSaved))
	if err != nil {
		return err
	}
	userSession.ReleaseState()
	return nil
}
