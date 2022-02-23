package repository

import (
	ecoerror "economicus/internal/error"
	"economicus/internal/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type QuantRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewQuantRepository(db *gorm.DB, logger *log.Logger) QuantRepositoryFactory {
	return &QuantRepository{
		db:     db,
		logger: logger,
	}
}

// GetAllQuants returns all uploaded quants
func (repo *QuantRepository) GetAllQuants(userID uint, option *models.QueryOption) (models.Quants, error) {
	var quants models.Quants

	sql := fmt.Sprintf("select * from quants as q "+
		"join profiles up on q.user_id = up.user_id "+
		"where q.name like '%%%s%%' or up.nickname like '%%%s%%' "+
		"order by q.user_id in (select following_id from followings where followings.user_id = %d) desc, %s "+
		"limit %d offset %d;",
		option.Query, option.Query, userID, option.Order, option.PerPage, option.Page*option.PerPage)

	if err := repo.db.Raw(sql).Find(&quants).Error; err != nil {
		repo.logger.Errorf("error in GetAllQuants: %v\n", err)
		return nil, err
	}

	return quants, nil
}

// GetFollowingsQuants returns quants of user's followings
func (repo *QuantRepository) GetFollowingsQuants(userID uint, option *models.QueryOption) (models.Quants, error) {
	var quants models.Quants

	err := repo.db.Raw("select * from quants "+
		"where user_id in (select following_id from user_followings "+
		"where followings.user_id = ?)"+
		"order by ? limit ? offset ?",
		userID, option.Order, option.PerPage, option.Page*option.PerPage).Find(&quants).Error

	if err != nil {
		repo.logger.Errorf("error in GetFollowingsQuants: %v\n", err)
		return nil, err
	}

	return quants, nil
}

// GetQuant returns a quant of quant id
func (repo *QuantRepository) GetQuant(quantID uint) (*models.Quant, error) {
	var quant models.Quant

	if err := repo.db.First(&quant, quantID).Error; err != nil {
		repo.logger.Errorf("error in GetQuant: %v\n", err)
		return nil, err
	}

	return &quant, nil
}

// GetMyQuants returns quants of the user
func (repo *QuantRepository) GetMyQuants(userID uint) (models.Quants, error) {
	var quants models.Quants

	if err := repo.db.Model(&models.Quant{}).Where("user_id = ?", userID).Find(&quants).Error; err != nil {
		repo.logger.Errorf("error in GetMyQuants: %v\n", err)
		return nil, err
	}

	return quants, nil
}

func (repo *QuantRepository) CheckQuantPermission(userID, quantID uint) error {
	quant, err := repo.GetQuant(quantID)
	if err != nil {
		repo.logger.Errorf("error in CheckQuantPermission: %v\n", err)
		return err
	}

	if quant.UserID != userID {
		return ecoerror.ErrPermissionDenied
	}

	return nil
}

func (repo *QuantRepository) CheckModelName(name string) error {
	err := repo.db.Where("name = ?", name).First(&models.Quant{}).Error
	if err == nil {
		return ecoerror.ErrDuplicateModelName
	} else if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		repo.logger.Errorf("error in CheckModelName: %v\n", err)
		return err
	}
}

// CreateQuant creates a quant
func (repo *QuantRepository) CreateQuant(quant *models.Quant) (uint, error) {
	if err := repo.db.Create(quant).Error; err != nil {
		repo.logger.Errorf("error in CreateQuant: %v\n", err)
		return 0, err
	}
	return quant.ID, nil
}

func (repo *QuantRepository) CreateQuantOption(quantOption *models.QuantOption) error {
	if err := repo.db.Create(quantOption).Error; err != nil {
		repo.logger.Errorf("error in CreateQuantOption: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepository) UpdateQuant(quantID uint, data map[string]interface{}) error {
	if err := repo.db.First(&models.Quant{}, quantID).Updates(data).Error; err != nil {
		repo.logger.Errorf("error in UpdateQuant: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepository) UpdateQuantOption(option *models.QuantOption) error {
	err := repo.db.Save(option).Error
	if err != nil {
		repo.logger.Errorf("error in UpdateQuantOption while finding quant")
		return err
	}

	return nil
}

func (repo *QuantRepository) DeleteQuant(quantID uint) error {
	if err := repo.db.Delete(&models.Quant{}, quantID).Error; err != nil {
		repo.logger.Errorf("error in DeleteQuant: %v\n", err)
		return err
	}
	return nil
}
