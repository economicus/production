package service

import (
	"economicus/internal/api/repository"
	"economicus/internal/models"
)

type CommentService struct {
	repo repository.CommentRepositoryFactory
}

func NewCommentService(repo repository.CommentRepositoryFactory) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) GetCommentsAndReplies(quantID uint) (models.Comments, error) {
	return s.repo.GetCommentsAndReplies(quantID)
}

func (s *CommentService) GetComment(commentID uint) (*models.Comment, error) {
	return s.repo.GetComment(commentID)
}

func (s *CommentService) CreateComment(userID, quantID uint, content string) error {
	comment := models.NewComment(userID, quantID, content)

	return s.repo.CreateComment(comment)
}

func (s *CommentService) UpdateComment(commentID uint, content string) error {
	return s.repo.UpdateComment(commentID, content)
}

func (s *CommentService) DeleteComment(commentID uint) error {
	return s.repo.DeleteComment(commentID)
}
