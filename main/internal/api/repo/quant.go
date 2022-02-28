package repo

import (
	"fmt"
	"gorm.io/gorm"
	e "main/internal/core/error"
	"main/internal/core/model"
	"main/internal/pkg/logger"
)

type QuantRepo struct {
	db *gorm.DB
}

func NewQuantRepo(db *gorm.DB) *QuantRepo {
	return &QuantRepo{
		db: db,
	}
}

// GetAllQuants returns all uploaded quants
func (repo *QuantRepo) GetAllQuants(userID uint, option *model.Query) (model.Quants, error) {
	var quants model.Quants

	sql := fmt.Sprintf("select * from quants as q "+
		"join profiles up on q.user_id = up.user_id "+
		"where q.name like '%%%s%%' or up.nickname like '%%%s%%' "+
		"order by q.user_id in (select following_id from followings where followings.user_id = %d) desc, %s "+
		"limit %d offset %d;",
		option.Keyword, option.Keyword, userID, option.Order, option.PerPage, option.Page*option.PerPage)

	if err := repo.db.Raw(sql).Find(&quants).Error; err != nil {
		logger.Logger.Errorf("error in GetAllQuants: %v\n", err)
		return nil, err
	}
	return quants, nil
}

// GetMyQuants returns quants of the user
func (repo *QuantRepo) GetMyQuants(userID uint) (model.Quants, error) {
	var quants model.Quants

	if err := repo.db.Model(&model.Quant{}).Where("user_id = ?", userID).Find(&quants).Error; err != nil {
		logger.Logger.Errorf("error in GetMyQuants: %v\n", err)
		return nil, err
	}
	return quants, nil
}

// GetQuant returns a quant of quant id
func (repo *QuantRepo) GetQuant(quantID uint) (*model.Quant, error) {
	var quant model.Quant

	if err := repo.db.First(&quant, quantID).Error; err != nil {
		logger.Logger.Errorf("error in GetQuant: %v\n", err)
		return nil, err
	}
	return &quant, nil
}

func (repo *QuantRepo) CheckModelName(name string) error {
	err := repo.db.Where("name = ?", name).First(&model.Quant{}).Error
	if err == nil {
		return e.ErrDuplicateModelName
	} else if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		logger.Logger.Errorf("error in CheckModelName: %v\n", err)
		return err
	}
}

// CreateQuant creates a quant
func (repo *QuantRepo) CreateQuant(quant *model.Quant) (uint, error) {
	if err := repo.db.Create(quant).Error; err != nil {
		logger.Logger.Errorf("error in CreateQuant: %v\n", err)
		return 0, err
	}
	return quant.ID, nil
}

func (repo *QuantRepo) CreateQuantOption(quantOption *model.QuantOption) error {
	if err := repo.db.Create(quantOption).Error; err != nil {
		logger.Logger.Errorf("error in CreateQuantOption: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepo) UpdateQuant(quantID uint, data map[string]interface{}) error {
	if err := repo.db.First(&model.Quant{}, quantID).Updates(data).Error; err != nil {
		logger.Logger.Errorf("error in UpdateQuant: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepo) UpdateQuantOption(quantID uint, req map[string]interface{}) error {
	err := repo.db.Where("quant_id = ?", quantID).First(&model.QuantOption{}).Updates(req).Error
	if err != nil {
		logger.Logger.Errorf("error in UpdateQuantOption while finding quant: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepo) DeleteQuant(quantID uint) error {
	if err := repo.db.Delete(&model.Quant{}, quantID).Error; err != nil {
		logger.Logger.Errorf("error in DeleteQuant: %v\n", err)
		return err
	}
	return nil
}
