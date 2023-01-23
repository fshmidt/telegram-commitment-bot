package telegram

import (
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func makeKeyBoard(commits []telegram_commitment_bot.CommitStruct) [][]tgbotapi.InlineKeyboardButton {

	var keyboard [][]tgbotapi.InlineKeyboardButton

	num := len(commits)
	numRows := num / 4

	if num%4 != 0 {
		numRows += 1
	}

	var ind int
	for r := 0; r < numRows; r++ {
		var row []tgbotapi.InlineKeyboardButton
		for k := 0; k < 4; k++ {

			var key tgbotapi.InlineKeyboardButton
			key.Text = strconv.Itoa(ind + 1)
			data := headerDelete + commits[ind].MakeUis()
			key.CallbackData = &data

			row = append(row, key)

			ind += 1
			if ind == num {
				break
			}
		}
		keyboard = append(keyboard, row)
		if ind == num {
			return keyboard
		}
	}
	return keyboard
}
