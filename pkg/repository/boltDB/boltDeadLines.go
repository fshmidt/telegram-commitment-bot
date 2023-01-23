package boltDB

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	"github.com/fshmidt/telegram-commitment-bot/pkg/config"
	"log"
	"time"
)

func (r *CommitRepository) CheckDl(messages config.Messages) ([]telegram_commitment_bot.CommitStruct, error) {

	DlCommits, err := r.CycleDlSearch(messages)
	if err != nil {
		return nil, JsonError
	}

	//updating new deadline info

	err = r.db.Batch(func(tx *bolt.Tx) error {
		for _, commit := range DlCommits {
			if err := updateCommit(tx, commit, UsersBucket, UpdateCommit); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, CantRewriteError
	}

	return DlCommits, err
}

func (r *CommitRepository) CycleDlSearch(messages config.Messages) ([]telegram_commitment_bot.CommitStruct, error) {

	var DlCommits []telegram_commitment_bot.CommitStruct

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))
		var num int = 1
		return b.ForEach(func(key, val []byte) error {

			data := b.Get(key)
			var commit telegram_commitment_bot.CommitStruct
			if err := json.Unmarshal(data, &commit); err != nil {
				return err
			}

			if commit.UserName != "" && r.DeadlineSoon(&commit, messages) {
				DlCommits = append(DlCommits, commit)
			}

			log.Println("CURRENT NUMBER OF ACTIVE PROMISES:", num, "User", commit.UserName, commit.UserID, "\nPromise:", commit.Commitment)
			num++
			return nil
		})
	})
	return DlCommits, err
}

func (r *CommitRepository) DeadlineSoon(c *telegram_commitment_bot.CommitStruct, messages config.Messages) bool {

	dl := c.Deadline.Sub(time.Now())

	if dl < 0 {
		r.Done(c.MakeUis(), c.ChatID, SuxBucket)
		//return false/////////////////////////////////////////////
	}

	var reminder string

	round := c.RoundDl(dl, messages)
	if round != "" {
		_, ok := c.RoundRemind[round]
		if ok == false {
			c.RoundRemind[round] = true
			reminder = round
		}
	}

	percents, scales := c.Percents(dl, messages)
	if scales != "" {
		c.Scale = scales
	}
	if percents != "" {
		_, ok := c.PercRemind[percents]
		if ok == false {
			c.PercRemind[percents] = true
			reminder += "\n" + c.Scale + "\n" + percents
		}
	} else if reminder != "" {
		reminder += "\n" + c.Scale
	}

	if reminder != "" {
		c.CurReminder = reminder
		return true
	}

	return false
}
