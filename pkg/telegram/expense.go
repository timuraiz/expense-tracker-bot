package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"github.com/timuraiz/expense-tracker-bot/pkg/telegram/session"
	"strconv"
	"strings"
	"time"
)

var (
	askMoneyCost = session.NewState("askMoneyCost")
	saveExpense  = session.NewState("saveInDb")
)

func handleAddExpenseCommand(b *Bot, message *tgbotapi.Message) error {

}

func handleAskCategory(b *Bot, message *tgbotapi.Message) error {
	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return err
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

	userSession.Data["category"] = message.Text

	if err != nil {
		return err
	}
	_, err = b.Bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.ExpenseMoneyCost))
	if err != nil {
		return err
	}
	userSession.SetState(&saveExpense)
	return nil
}

func handleSaveExpense(b *Bot, message *tgbotapi.Message) error {
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
	_, err = b.Db.AddExpense(expenseDetail)
	if err != nil {
		return unableToSaveExpenseError
	}
	_, err = b.Bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.ExpenseSaved))

	return err
}
