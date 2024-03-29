package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/telegram/session"
)

var (
	first = session.NewState("first")
)

func handleStartCommand(b *Bot, message *tgbotapi.Message) error {

	userSession, err := b.SessionStorage.GetSession(message.Chat.ID)
	if err != nil {
		return unableToReachUserSessionError
	}
	switch userSession.State.GetName() {
	case first.GetName():
		msg := tgbotapi.NewMessage(message.Chat.ID, "You?? Again??")
		_, err = b.Bot.Send(msg)
		userSession.SetState(&session.NullState)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.Start)
		_, err = b.Bot.Send(msg)
		userSession.SetState(&first)

	}
	return err
}
