package telegram

import (
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

const (
	commandStart     = "start"
	commandCommit    = "promise"
	commandMyCommits = "mypromises"
	commandDelete    = "delete"
	headerDelete     = "DELETE"
	keepWorkingTail  = "Keep Working"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	if promiseInside(message.Text) {
		msg.Text = b.messages.Responses.Initiate
		_, err := b.bot.Send(msg)
		return err
	}

	return nil
}

func promiseInside(message string) bool {
	for _, phrases := range telegram_commitment_bot.Promices {
		if strings.Contains(message, phrases) {
			return true
		}
	}
	return false
}

func (b *Bot) handleDone(query *tgbotapi.CallbackQuery) error {

	userId := strconv.FormatInt(query.From.ID, 10)

	if userId != query.Data[:len(userId)] { //validation of user Id
		return nil
	}

	checkCommit := b.commitRepository.GetByUis(query.Data, "done_bucket")
	if checkCommit.Commitment != "" {
		return b.removeButtons(query.Message)
	}

	checkCommit = b.commitRepository.GetByUis(query.Data, "sux_bucket")
	if checkCommit.Commitment != "" {
		return b.removeButtons(query.Message)
	}

	if err := b.commitRepository.Done(query.Data, query.Message.Chat.ID, "done_bucket"); err != nil {
		return err
	}

	if err := b.removeButtons(query.Message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(query.Message.Chat.ID, b.messages.Responses.Congrats)
	msg.Text = b.messages.Responses.Congrats
	msg.ReplyToMessageID = query.Message.MessageID

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleDeleteKey(query *tgbotapi.CallbackQuery) error {

	userId := strconv.FormatInt(query.From.ID, 10)
	offset := len(headerDelete)

	if userId != query.Data[offset:offset+len(userId)] { //validation of user Id
		return nil
	}

	uis := query.Data[offset:]

	if err := b.commitRepository.Delete(uis, "done_bucket"); err != nil {
		return err
	}

	if err := b.removeButtons(query.Message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(query.Message.Chat.ID, b.messages.Responses.CommitIsDeleted)
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) removeButtons(message *tgbotapi.Message) error {

	cfg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, message.Text)

	_, err := b.bot.Send(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleKeyOk(query *tgbotapi.CallbackQuery) {

	userId := strconv.FormatInt(query.From.ID, 10)

	if userId != query.Data[:len(userId)] { //validate user id
		return
	}

	message := query.Message

	cfg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, message.Text) ////////
	_, _ = b.bot.Send(cfg)
}
