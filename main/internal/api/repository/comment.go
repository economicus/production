package repository

import (
	"economicus/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewCommentRepository(db *gorm.DB, logger *log.Logger) CommentRepositoryFactory {
	return &CommentRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *CommentRepository) GetCommentsAndReplies(quantID uint) (models.Comments, error) {
	var comments models.Comments
	err := repo.db.Preload("Replies").Where("quant_id = ?", quantID).Find(&comments).Error
	if err != nil {
		repo.logger.Errorf("error in GetCommentsAndReplies: %v\n", err)
		return nil, err
	}
	return comments, nil
}

func (repo *CommentRepository) GetComment(commentID uint) (*models.Comment, error) {
	var comment models.Comment
	if err := repo.db.First(&comment, commentID).Error; err != nil {
		repo.logger.Errorf("error in GetComment: %v\n", err)
		return nil, err
	}
	return &comment, nil
}

func (repo *CommentRepository) CreateComment(comment *models.Comment) error {
	if err := repo.db.Create(comment).Error; err != nil {
		repo.logger.Errorf("error in CreateComment: %v\n", err)
		return err
	}
	return nil
}

func (repo *CommentRepository) UpdateComment(commentID uint, content string) error {
	if err := repo.db.First(&models.Comment{}, commentID).Update("content", content).Error; err != nil {
		repo.logger.Errorf("error in UpdateComment: %v\n", err)
		return err
	}
	return nil
}

func (repo *CommentRepository) DeleteComment(commentID uint) error {
	if err := repo.db.Delete(&models.Comment{}, commentID).Error; err != nil {
		repo.logger.Errorf("error in DeleteComment: %v\n", err)
		return err
	}
	return nil
}
