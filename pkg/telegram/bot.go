package telegram

import (
	"github.com/fshmidt/telegram-commitment-bot/pkg/config"
	"github.com/fshmidt/telegram-commitment-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type Bot struct {
	bot              *tgbotapi.BotAPI
	commitRepository repository.CommitRepository
	messages         config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, cr repository.CommitRepository, messages config.Messages) *Bot {
	return &Bot{bot: bot, commitRepository: cr, messages: messages}
}

func (b *Bot) Start() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	b.periodCheck()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		// If we got a message

		if update.Message != nil {

			if update.Message.IsCommand() {

				err := b.handleCommand(update.Message)
				if err != nil {
					b.handleError(update.Message, err)
				}
				continue
			}

			err := b.handleMessage(update.Message)
			if err != nil {
				b.handleError(update.Message, err)
			}

			// key update

		} else if update.CallbackQuery != nil {

			data := update.CallbackQuery.Data

			if update.CallbackQuery.Data[:6] == headerDelete {

				err := b.handleDeleteKey(update.CallbackQuery)
				if err != nil {
					b.handleKeyError(update.CallbackQuery, err)
				}

			} else if data[len(data)-len(keepWorkingTail):] != keepWorkingTail {

				err := b.handleDone(update.CallbackQuery)
				if err != nil {
					b.handleKeyError(update.CallbackQuery, err)
				}

			} else {

				b.handleKeyOk(update.CallbackQuery)

			}
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) periodCheck() {
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				b.checkDeadlines()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
