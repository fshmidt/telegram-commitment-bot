package main

import (
	"github.com/boltdb/bolt"
	"github.com/fshmidt/telegram-commitment-bot/pkg/config"
	"github.com/fshmidt/telegram-commitment-bot/pkg/repository/boltDB"
	"github.com/fshmidt/telegram-commitment-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	commitRepository := boltDB.NewCommitRepository(db)

	telegramBot := telegram.NewBot(bot, commitRepository, cfg.Messages)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.BoltDb, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltDB.UsersBucket))
		return err
	}); err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltDB.DoneBucket))
		return err
	}); err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltDB.SuxBucket))
		return err
	}); err != nil {
		return nil, err
	}

	return db, err
}
