package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/config"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
	Db  storage.Crud
	Cfg *config.Config
}

var UserStates map[int64]*State

func NewBot(bot *tgbotapi.BotAPI, db storage.Crud, cfg *config.Config) *Bot {
	UserStates = make(map[int64]*State)
	return &Bot{
		Bot: bot,
		Db:  db,
		Cfg: cfg,
	}
}

type State string

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		chatId := update.Message.Chat.ID
		var state *State

		state = UserStates[chatId]

		// Handle commands
		if update.Message.IsCommand() {
			state, err := b.handleCommand(update.Message, state)
			UserStates[chatId] = state
			if err != nil {
				b.handleError(chatId, err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
