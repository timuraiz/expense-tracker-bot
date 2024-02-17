package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timuraiz/expense-tracker-bot/pkg/config"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage"
	"github.com/timuraiz/expense-tracker-bot/pkg/storage/postgre"
	"github.com/timuraiz/expense-tracker-bot/pkg/telegram"
	"github.com/timuraiz/expense-tracker-bot/pkg/telegram/session"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = true

	db, err := storage.SetupDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := postgre.NewPostgresRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	sessionStorage := session.NewInMemorySessionStorage()
	bot := telegram.NewBot(botApi, repo, cfg, sessionStorage)
	if err := bot.Start(); err != nil {
		log.Panic(err)
	}
}
