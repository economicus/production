package repository

import (
	"economicus/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReplyRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewReplyRepository(db *gorm.DB, logger *log.Logger) ReplyRepositoryFactory {
	return &ReplyRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *ReplyRepository) GetReply(replyID uint) (*models.Reply, error) {
	var reply models.Reply
	err := repo.db.First(&reply, replyID).Error
	return &reply, err
}

func (repo *ReplyRepository) CreateReply(reply *models.Reply) error {
	if err := repo.db.Create(reply).Error; err != nil {
		repo.logger.Errorf("error in CreateReply: %v\n", err)
		return err
	}
	return nil
}

func (repo *ReplyRepository) UpdateReply(replyID uint, content string) error {
	if err := repo.db.First(&models.Reply{}, replyID).Update("content", content).Error; err != nil {
		repo.logger.Errorf("error in UpdateReply: %v\n", err)
		return err
	}
	return nil
}

func (repo *ReplyRepository) DeleteReply(replyID uint) error {
	if err := repo.db.Delete(&models.Reply{}, replyID).Error; err != nil {
		repo.logger.Errorf("error in DeleteReply: %v\n", err)
		return err
	}
	return nil
}
