package service

import (
	"economicus/internal/api/repository"
	"economicus/internal/models"
)

type ReplyService struct {
	repo repository.ReplyRepositoryFactory
}

func NewReplyService(repo repository.ReplyRepositoryFactory) *ReplyService {
	return &ReplyService{
		repo: repo,
	}
}

func (s *ReplyService) GetReply(replyID uint) (*models.Reply, error) {
	return s.repo.GetReply(replyID)
}

func (s *ReplyService) CreateReply(userID, commentID uint, content string) error {
	reply := models.NewReply(userID, commentID, content)

	return s.repo.CreateReply(reply)
}

func (s *ReplyService) UpdateReply(replyID uint, content string) error {
	return s.repo.UpdateReply(replyID, content)
}

func (s *ReplyService) DeleteReply(replyID uint) error {
	return s.repo.DeleteReply(replyID)
}
