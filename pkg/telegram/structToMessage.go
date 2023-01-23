package telegram

import (
	"fmt"
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
)

func commitsInMsg(commits []telegram_commitment_bot.CommitStruct) string {

	if len(commits) == 0 {
		return ""
	}

	var msg string
	var doneInd, suxInd int

	for ind, commit := range commits {

		if commit.Ok == true {

			if msg == "" {
				msg = "ТВОИ ТЕКУЩИЕ ОБЕЩАНИЯ:\n▪▪▪▪▪▪▪▪▪▪▪▪▪\n"
			}
			msg += timedCommit(commit, ind)

		} else {

			var number int
			if commit.Bucket == "done_bucket" {
				number = doneInd
				doneInd += 1
			} else {
				number = suxInd
				suxInd += 1
			}
			msg += timedCommit(commit, number)
		}
	}
	return msg
}

func timedCommit(commit telegram_commitment_bot.CommitStruct, ind int) string {

	if ind == -1 {
		return fmt.Sprintf("%s до %s.\n", commit.Commitment, commit.Deadline.Format("02.01.2006 15:04:05")[:10])
	}

	if commit.Ok == false {
		if commit.Bucket == "done_bucket" {

			var doneBegin string
			if ind == 0 {
				doneBegin = "\n😎УЖЕ СДЕЛАНО😎\n-----------------------------------------\n"
			}
			return fmt.Sprintf("%s%d. %s ✔️\n-----------------------------------------\n", doneBegin, ind+1, commit.Commitment)
		} else {

			var suxBegin string
			if ind == 0 {
				suxBegin = "\n💩ПРОСОСАНО💩\n-----------------------------------------\n"
			}
			return fmt.Sprintf("%s%d. %s ️☠️\n-----------------------------------------\n", suxBegin, ind+1, commit.Commitment)
		}
	}
	return fmt.Sprintf("%d. %s до %s.\n%s\n▪▪▪▪▪▪▪▪▪▪▪▪▪\n", ind+1, commit.Commitment, commit.Deadline.Format("02.01.2006 15:04:05")[:10], commit.Scale)
}
