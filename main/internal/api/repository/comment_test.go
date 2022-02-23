package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentRepository_GetCommentsAndReplies(t *testing.T) {
	a := assert.New(t)
	comments, err := commentRepo.GetCommentsAndReplies(uint(1))
	a.NoError(err, err)
	a.NotEmpty(comments, "comments should not be empty")
}

func TestCommentRepository_CreateComment(t *testing.T) {
	a := assert.New(t)
	err := commentRepo.CreateComment(uint(1), uint(2), "comment test")
	a.NoError(err, err)
}

func TestCommentRepository_UpdateComment(t *testing.T) {
	a := assert.New(t)
	err := commentRepo.UpdateComment(uint(1), "comment edit test")
	a.NoError(err, err)
}

func TestCommentRepository_DeleteComment(t *testing.T) {
	a := assert.New(t)
	err := commentRepo.DeleteComment(uint(1))
	a.NoError(err, err)
}
