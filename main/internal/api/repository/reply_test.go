package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplyRepository_CreateReply(t *testing.T) {
	a := assert.New(t)
	err := replyRepo.CreateReply(uint(1), uint(2), "create test")
	a.NoError(err, err)
}

func TestReplyRepository_UpdateReply(t *testing.T) {
	a := assert.New(t)
	err := replyRepo.UpdateReply(uint(2), "create test")
	a.NoError(err, err)
}

func TestReplyRepository_DeleteReply(t *testing.T) {
	a := assert.New(t)
	err := replyRepo.DeleteReply(uint(1))
	a.NoError(err, err)
}
