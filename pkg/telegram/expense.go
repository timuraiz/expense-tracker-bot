package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"strconv"
	"strings"
	"time"
)

func handleAddExpenseCommand(b *Bot, message *tgbotapi.Message, state *State) error {
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
