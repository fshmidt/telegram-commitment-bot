package boltDB

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	"github.com/fshmidt/telegram-commitment-bot/pkg/repository"
	"strconv"
)

func (r *CommitRepository) GetCurrentCommitNum(userId int64) (int, error) {

	var num int

	strId := strconv.FormatInt(userId, 10)
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))

		err := b.ForEach(func(key, val []byte) error {

			if string(key)[:len(strId)] == strId {
				num++
			}
			return nil
		})
		return err
	})
	return num, err
}

func (r *CommitRepository) Get(userId int64, chatId int64) ([]telegram_commitment_bot.CommitStruct, error) {

	var commits []telegram_commitment_bot.CommitStruct
	strId := strconv.FormatInt(userId, 10)

	err := r.GetUserCommits(&commits, strId, chatId, UsersBucket)
	if err != nil {
		return nil, JsonError
	}
	err = r.GetUserCommits(&commits, strId, chatId, DoneBucket)
	if err != nil {
		return nil, JsonError
	}
	err = r.GetUserCommits(&commits, strId, chatId, SuxBucket)
	if err != nil {
		return nil, JsonError
	}
	return commits, err
}

func (r *CommitRepository) GetUserCommits(commits *[]telegram_commitment_bot.CommitStruct, id string, chatId int64, bucket repository.Bucket) error {
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		err := b.ForEach(func(key, val []byte) error {

			if string(key)[:len(id)] == string(id) {
				data := b.Get(key)
				var commit telegram_commitment_bot.CommitStruct
				if err := json.Unmarshal(data, &commit); err != nil {
					return err
				}
				commit.Bucket = string(bucket)
				if commit.ChatID == chatId {
					*commits = append(*commits, commit)
				}
			}
			return nil
		})
		return err
	})
	return err
}

func (r *CommitRepository) GetByUis(uis string, bucket repository.Bucket) telegram_commitment_bot.CommitStruct {

	var commit telegram_commitment_bot.CommitStruct
	r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		err := b.ForEach(func(key, val []byte) error {

			if string(key) == uis {
				data := b.Get(key)
				if err := json.Unmarshal(data, &commit); err != nil {
					return err
				} else {
					return nil
				}
			}
			return nil
		})
		return err
	})
	return commit
}
