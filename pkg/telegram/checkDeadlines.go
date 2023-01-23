package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"time"
)

func (b *Bot) checkDeadlines() error {

	commits, err := b.commitRepository.CheckDl(b.messages)
	if err != nil {
		return err
	}

	var keepWorking []string = []string{b.messages.Responses.Ok1, b.messages.Responses.Ok2, b.messages.Responses.Ok3, b.messages.Responses.Ok4}

	for _, commit := range commits {

		rand.Seed(time.Now().UnixNano())

		var keyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(b.messages.Responses.Done, commit.MakeUis()),
				tgbotapi.NewInlineKeyboardButtonData(keepWorking[rand.Intn(4)], commit.MakeUis()+keepWorkingTail),
			),
		)

		finalReminder := fmt.Sprintf("@%s, %s\nОбещание: %s", commit.UserName, commit.CurReminder, timedCommit(commit, -1))
		msg := tgbotapi.NewMessage(commit.ChatID, finalReminder)
		msg.ReplyMarkup = keyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
