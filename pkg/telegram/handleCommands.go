package telegram

import (
	"fmt"
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	log.Printf("COMMAND [%s] %s", message.From.UserName, message.Text)

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandCommit:
		return b.handleCommitCommand(message)
	case commandMyCommits:
		return b.handleGetCommitsCommand(message)
	case commandDelete:
		return b.handleDeleteCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.Text = fmt.Sprintf(b.messages.Responses.Start)
	msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleCommitCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.Text = b.messages.Responses.Commit
	msg.ReplyToMessageID = message.MessageID

	num, err := b.commitRepository.GetCurrentCommitNum(message.From.ID)

	if err != nil {
		return err
	}

	if num >= 15 {
		msg.Text = b.messages.Responses.TooManyCommits
		_, err := b.bot.Send(msg)
		return err
	}

	commit, err := mesToCom(message)

	if err != nil {
		return err
	}
	commit.Scale = b.messages.Scales.Start

	err = b.commitRepository.Save(commit, "users_bucket")
	if err != nil {
		return err
	}

	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleGetCommitsCommand(message *tgbotapi.Message) error {

	if len(message.Text) > 150 {
		return TooLongError
	}
	commits, err := b.commitRepository.Get(message.From.ID, message.Chat.ID) //

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.ReplyToMessageID = message.MessageID
	msg.Text = commitsInMsg(commits)

	if msg.Text == "" {
		msg.Text = b.messages.Responses.NoCommits
	}
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.Text = b.messages.Responses.Unknown
	msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleDeleteCommand(message *tgbotapi.Message) error {

	var commits []telegram_commitment_bot.CommitStruct
	strId := strconv.FormatInt(message.From.ID, 10)

	b.commitRepository.GetUserCommits(&commits, strId, message.Chat.ID, "done_bucket")
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.DeleteHeader)

	if len(commits) == 0 {
		msg.Text += b.messages.Responses.NoDoneCommits
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		return nil
	}
	msg.Text += commitsInMsg(commits)

	var keyboard tgbotapi.InlineKeyboardMarkup
	keyboard.InlineKeyboard = makeKeyBoard(commits)

	msg.ReplyMarkup = keyboard

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
