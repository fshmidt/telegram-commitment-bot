package boltDB

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	telegram_commitment_bot "github.com/fshmidt/telegram-commitment-bot"
	"github.com/fshmidt/telegram-commitment-bot/pkg/repository"
	"strconv"
	"strings"
)

const (
	UsersBucket  repository.Bucket = "users_bucket"
	DoneBucket   repository.Bucket = "done_bucket"
	SuxBucket    repository.Bucket = "sux_bucket"
	CreateCommit                   = 1
	UpdateCommit                   = 2
)

var (
	ExistingCommitError = errors.New("commit exists already")
	JsonError           = errors.New("marshal/unmarshal problem")
	CantRewriteError    = errors.New("can't rewrite")
)

type CommitRepository struct {
	db *bolt.DB
}

func NewCommitRepository(db *bolt.DB) *CommitRepository {
	return &CommitRepository{db: db}
}

func (r *CommitRepository) Save(commit telegram_commitment_bot.CommitStruct, bucket repository.Bucket) error {
	commit.Bucket = string(bucket)
	fmt.Println("IN SAVE COMMIT>USERNAME", commit.UserName)
	return r.db.Update(func(tx *bolt.Tx) error {
		return updateCommit(tx, commit, bucket, CreateCommit)
	})
}

func updateCommit(tx *bolt.Tx, commit telegram_commitment_bot.CommitStruct, bucket repository.Bucket, flag int) error {

	uis := commit.MakeUis()
	byteUis := []byte(uis)
	saveToBd := tx.Bucket([]byte(bucket))

	if notExist := saveToBd.Get(byteUis); notExist != nil && flag == CreateCommit {
		return ExistingCommitError
	}

	if buf, err := json.Marshal(commit); err != nil {
		return err
	} else if err := saveToBd.Put(byteUis, buf); err != nil {
		return err
	}
	return nil
}

func (r *CommitRepository) Done(uis string, chatId int64, destBucket repository.Bucket) error {

	err := r.Delete(uis, UsersBucket)
	if err != nil {
		return err
	}

	var commit telegram_commitment_bot.CommitStruct
	commit.UserID, err = strconv.ParseUint(strings.Split(uis, "*")[0], 10, 64)
	commit.Commitment = strings.Split(uis, "*")[1]
	commit.ChatID = chatId

	err = r.Save(commit, destBucket)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommitRepository) Delete(uis string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		transaction := tx.Bucket([]byte(bucket))

		// Marshal and save the encoded user.
		if err := transaction.Delete([]byte(uis)); err != nil {
			return err
		}
		return nil
	})
}
