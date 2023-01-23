package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	noCommitError       = errors.New("there's no commitment in message")
	noDateError         = errors.New("there's bad data in message")
	badYearError        = errors.New("something wrong with in your Year commitment")
	badMonthError       = errors.New("something wrong with Month in your commitment")
	badDayError         = errors.New("something wrong with Day in your commitment")
	ExistingCommitError = errors.New("commit exists already")
	JsonError           = errors.New("marshal/unmarshal problem")
	tooSoonError        = errors.New("date is in past or today")
	TooLongError        = errors.New("too long promise")
)

func (b *Bot) handleError(message *tgbotapi.Message, err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная ошибка")
	msg.ReplyToMessageID = message.MessageID

	switch err {
	case noCommitError:
		msg.Text = b.messages.Errors.NoCommit
	case noDateError:
		msg.Text = b.messages.Errors.NoDate
	case badYearError:
		msg.Text = b.messages.Errors.BadYear
	case badMonthError:
		msg.Text = b.messages.Errors.BadMonth
	case badDayError:
		msg.Text = b.messages.Errors.BadDay
	case ExistingCommitError:
		msg.Text = b.messages.Errors.ExistingCommit
	case JsonError:
		msg.Text = b.messages.Errors.JsonError
	case tooSoonError:
		msg.Text = b.messages.Errors.TooSoon
	case TooLongError:
		msg.Text = b.messages.Errors.TooLongError
	}
	_, err = b.bot.Send(msg)
}

func (b *Bot) handleKeyError(query *tgbotapi.CallbackQuery, err error) {

	msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Неизвестная ошибка")
	msg.ReplyToMessageID = query.Message.MessageID

	switch err {

	case ExistingCommitError:
		msg.Text = b.messages.Errors.ExistingCommit
	case JsonError:
		msg.Text = b.messages.Errors.JsonError
	}
	_, err = b.bot.Send(msg)
}
