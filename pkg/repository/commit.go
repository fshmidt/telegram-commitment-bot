package repository

import (
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	"github.com/fshmidt/telegram-commitment-bot/pkg/config"
)

type Bucket string

type CommitRepository interface {
	Save(commit telegram_commitment_bot.CommitStruct, bucket Bucket) error
	Get(userId int64, chatId int64) ([]telegram_commitment_bot.CommitStruct, error)
	GetByUis(uis string, bucket Bucket) telegram_commitment_bot.CommitStruct
	GetCurrentCommitNum(userId int64) (int, error)
	GetUserCommits(c *[]telegram_commitment_bot.CommitStruct, id string, chatId int64, bucket Bucket) error
	Delete(uis string, bucket Bucket) error
	Done(uis string, chatId int64, destBucket Bucket) error
	CheckDl(messages config.Messages) ([]telegram_commitment_bot.CommitStruct, error)
	CycleDlSearch(messages config.Messages) ([]telegram_commitment_bot.CommitStruct, error)
}
