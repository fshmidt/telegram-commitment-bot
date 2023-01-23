package telegram

import (
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"time"
)

func mesToCom(message *tgbotapi.Message) (telegram_commitment_bot.CommitStruct, error) {
	var commit telegram_commitment_bot.CommitStruct

	deadline, err := parceCommitment(&commit, message)
	if err != nil {
		return commit, err
	}
	err = parceDate(&commit, deadline)
	if err != nil {
		return commit, err
	}

	if err != nil {
		return commit, err
	} else {
		commit.UserID = uint64(message.From.ID)
		commit.UserName = message.From.UserName
		commit.ChatID = message.Chat.ID
		commit.Created = time.Now()
		commit.Ok = true
		commit.PercRemind = make(map[string]bool)
		commit.RoundRemind = make(map[string]bool)
	}
	return commit, nil
}

func parceCommitment(commit *telegram_commitment_bot.CommitStruct, message *tgbotapi.Message) (string, error) {
	var deadline string
	var do int

	msg := message.Text
	msgToSlice := strings.Split(msg, " ")

	for i, word := range msgToSlice {
		if word != "/promise" && word != "#до" && word != "#ДО" && word != "#До" && do == 0 {
			commit.Commitment += " " + word
		} else if do == 0 {
			do = i
		}
		if i > do && do > 0 {
			deadline += word
		}
	}
	if do == 0 || commit.Commitment == "" {
		return "", noCommitError
	}
	return deadline, nil
}

func parceDate(commit *telegram_commitment_bot.CommitStruct, deadline string) (err error) {
	dlToSlice := strings.Split(deadline, ".")

	if len(dlToSlice) != 3 {
		return noDateError
	}

	month, err := strconv.Atoi(dlToSlice[1])
	if err != nil || month < 1 || month > 12 {
		return badMonthError
	}
	day, err := strconv.Atoi(dlToSlice[0])
	if err != nil || day < 1 || day > 31 {
		return badDayError
	}

	year, err := strconv.Atoi(dlToSlice[2])
	if err != nil || year < time.Now().Year() {
		return badYearError
	}

	parsedDl := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if parsedDl.Unix() <= time.Now().Unix() {
		return tooSoonError
	}
	commit.Deadline = parsedDl

	return nil
}
