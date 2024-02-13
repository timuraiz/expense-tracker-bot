package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	firstAttempt State = "firstAttempt"
)

func handleStartCommand(b *Bot, message *tgbotapi.Message, state *State) (*State, error) {
	var err error
	switch {
	case state == nil:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.Start)
		_, err = b.Bot.Send(msg)
		state = &firstAttempt
	case *state == firstAttempt:
		msg := tgbotapi.NewMessage(message.Chat.ID, "You have already called this command")
		_, err = b.Bot.Send(msg)
		state = nil
	}
	//msg := tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.Start)
	//_, err := b.Bot.Send(msg)
	return state, err
}
