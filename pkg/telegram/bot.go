package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/config"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"github.com/timuraiz/expense-tracker-bot/pkg/telegram/session"
)

type Bot struct {
	Bot            *tgbotapi.BotAPI
	Db             storage.Crud
	Cfg            *config.Config
	SessionStorage session.Session
}

func NewBot(bot *tgbotapi.BotAPI, db storage.Crud, cfg *config.Config, session session.Session) *Bot {
	return &Bot{
		Bot:            bot,
		Db:             db,
		Cfg:            cfg,
		SessionStorage: session,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if err := b.handleCommand(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
