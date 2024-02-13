package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUnknownCommand(b *Bot, message *tgbotapi.Message, state *State) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.Cfg.Responses.UnknownCommand)
	_, err := b.Bot.Send(msg)
	return err
}
