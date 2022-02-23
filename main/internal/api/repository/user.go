package repository

import (
	"economicus/commons"
	"economicus/internal/drivers"
	"economicus/internal/error"
	"economicus/internal/models"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db     *gorm.DB
	aws    *drivers.AWS
	logger *log.Logger
}

func NewUserRepository(db *gorm.DB, aws *drivers.AWS, logger *log.Logger) UserRepositoryFactory {
	return &UserRepository{
		db:     db,
		aws:    aws,
		logger: logger,
	}
}

// GetUsers return users
func (repo *UserRepository) GetUsers(option *models.QueryOption) (models.Users, error) {
	var users models.Users

	sql := fmt.Sprintf("select * "+
		"from users as u join profiles as p on u.id = p.user_id "+
		"where u.name like '%%%s%%' or p.nickname like '%%%s%%' "+
		"order by %s limit %d offset %d;",
		option.Query, option.Query, option.Order, option.PerPage, option.Page*option.PerPage)

	err := repo.db.Preload("Profile").Raw(sql).Find(&users).Error
	if err != nil {
		repo.logger.Errorf("error in GetUsers: %v\n", err)
		return nil, err
	}

	return users, nil
}

// GetUserByID finds user with user id
func (repo *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User

	err := repo.db.Preload("Profile").First(&user, userID).Error
	if err != nil {
		repo.logger.Errorf("error in GetUserByID: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) CheckNickname(nickname string) error {
	err := repo.db.Where("nickname = ?", nickname).First(&models.Profile{}).Error
	if err == nil {
		return ecoerror.ErrDuplicateNickname
	} else if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		repo.logger.Errorf("error in CheckNickname: %v\n", err)
		return err
	}
}

func (repo *UserRepository) CreateUser(email, password, name string) (uint, error) {
	user := models.NewUser(email, password, name)

	if err := repo.db.Create(user).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return 0, ecoerror.ErrDuplicateEmail
		} else {
			repo.logger.Errorf("error in CreateUser: %v\n", err)
			return 0, err
		}
	}

	return user.ID, nil
}

func (repo *UserRepository) CreateProfile(userID uint, nickname string, birth time.Time) error {
	profile := models.NewProfile(userID, nickname, birth)

	if err := repo.db.Create(profile).Error; err != nil {
		repo.logger.Errorf("error in CreateProfile: %v\n", err)
		return err
	}

	return nil
}

// UpdateUserProfile updates user's profile
func (repo *UserRepository) UpdateUserProfile(userID uint, data map[string]interface{}) error {
	var profile models.Profile

	err := repo.db.Where("user_id = ?", userID).First(&profile).Updates(data).Error
	if err != nil {
		repo.logger.Errorf("error in UpdateUserProfile: %v\n", err)
		return err
	}

	return nil
}

// UploadUserProfileImage uploads user's profile image path from s3
func (repo *UserRepository) UploadUserProfileImage(userID uint, filepath string) error {
	var profile models.Profile

	err := repo.db.Model(&profile).Where("user_id = ?", userID).Update("profile_image", filepath).Error
	if err != nil {
		repo.logger.Errorf("error in UploadUserProfileImage: %v\n", err)
		return err
	}

	return nil
}

// DeleteUser soft-deletes a user
func (repo *UserRepository) DeleteUser(ID uint) error {
	if err := repo.db.First(&models.User{}, ID).Update("user_active", 0).Error; err != nil {
		repo.logger.Errorf("error in DeleteUser while inactivating user: %v\n", err)
		return err
	}

	if err := repo.db.Delete(&models.User{}, ID).Error; err != nil {
		repo.logger.Errorf("error in DeleteUser while deleting user: %v\n", err)
		return err
	}

	return nil
}

// UpdatePassword updates a user's password
func (repo *UserRepository) UpdatePassword(userID uint, newPassword string) error {
	newHashedPwd, err := commons.HashPassword([]byte(newPassword))

	if err != nil {
		repo.logger.Errorf("error in UpdatePassword while hashing password: %v\n", err)
		return err
	}

	err = repo.db.First(&models.User{}, userID).Update("password", newHashedPwd).Error
	if err != nil {
		repo.logger.Errorf("error in UpdatePassword while updating password: %v\n", err)
		return err
	}

	return nil
}

// GetFollowings returns users who are followed by the user
func (repo *UserRepository) GetFollowings(userID uint) (models.Users, error) {
	var followings models.Users

	user, err := repo.GetUserByID(userID)
	if err != nil {
		repo.logger.Errorf("error in GetFollowings while finding user: %v\n", err)
		return nil, err
	}

	err = repo.db.Preload("Profile").Model(&user).Association("Followings").Find(&followings)
	if err != nil {
		repo.logger.Errorf("error in GetFollowings while following user: %v\n", err)
		return nil, err
	}

	return followings, err
}

// GetFollowers returns users who follow the user
func (repo *UserRepository) GetFollowers(userID uint) (models.Users, error) {
	var users models.Users

	sql := fmt.Sprintf("select * from users where id in (select user_id from followings where following_id = %d)", userID)
	if err := repo.db.Preload("Profile").Raw(sql).Find(&users).Error; err != nil {
		repo.logger.Errorf("error in GetFollowers: %v\n", err)
		return nil, err
	}

	return users, nil
}

// Follow adds to follower list of the user
func (repo *UserRepository) Follow(userID, followingID uint) error {
	sql := fmt.Sprintf("insert into followings (user_id, following_id) values (%d, %d)", userID, followingID)

	if err := repo.db.Exec(sql).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return ecoerror.ErrDuplicateData
		} else {
			repo.logger.Errorf("error in Follow: %v", err)
			return err
		}
	}

	return nil
}

// UnFollow deletes a following from follower list of the user
func (repo *UserRepository) UnFollow(userID, followingID uint) error {
	sql := fmt.Sprintf("delete from followings where user_id = %d and following_id = %d", userID, followingID)

	if err := repo.db.Exec(sql).Error; err != nil {
		repo.logger.Errorf("error in UnFollow: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepository) GetFavoriteQuants(userID uint) ([]*models.Quant, error) {
	var user models.User

	if err := repo.db.Preload("FavoriteQuants").First(&user, userID).Error; err != nil {
		repo.logger.Errorf("error in GetFavoriteQuants: %v", err)
		return nil, err
	}

	return user.FavoriteQuants, nil
}

func (repo *UserRepository) AddToFavoriteQuants(userID, quantID uint) error {
	sql := fmt.Sprintf("insert into user_favorite_quants (user_id, quant_id) values (%d, %d)", userID, quantID)

	if err := repo.db.Exec(sql).Error; err != nil {
		repo.logger.Errorf("error in AddToFavoriteQuants: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepository) DeleteFromFavoriteQuants(userID, quantID uint) error {
	sql := fmt.Sprintf("delete from user_favorite_quants where (user_id, quant_id) = (%d, %d)", userID, quantID)

	if err := repo.db.Exec(sql).Error; err != nil {
		repo.logger.Errorf("error in DeleteFromFavoriteQuants: %v", err)
		return err
	}

	return nil
}
