package service

import (
	"economicus/internal/api/repository"
	"economicus/internal/drivers"
	"economicus/internal/models"
	"fmt"
	"mime/multipart"
)

type UserService struct {
	repo repository.UserRepositoryFactory
	aws  *drivers.AWS
}

func NewUserService(repo repository.UserRepositoryFactory, aws *drivers.AWS) *UserService {
	return &UserService{
		repo: repo,
		aws:  aws,
	}
}

func (s *UserService) GetUsers(option *models.QueryOption) (models.Users, error) {
	return s.repo.GetUsers(option)
}

func (s *UserService) GetUserDataWithFields(userID uint, fields []string) (map[string]interface{}, error) {
	res := map[string]interface{}{}

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.Profile.ProfileImage = s.aws.GetFilePath(user.Profile.ProfileImage)

	data := user.ConvToMap()
	if len(fields) == 0 {
		return data, nil
	}
	for _, val := range fields {
		if _, ok := data[val]; ok {
			res[val] = data[val]
		}
	}

	return res, nil
}

func (s *UserService) Register(request *models.RegisterRequest) error {
	if err := s.repo.CheckNickname(request.Nickname); err != nil {
		return err
	}
	userID, err := s.repo.CreateUser(request.Email, request.Password, request.Name)
	if err != nil {
		return err
	}
	return s.repo.CreateProfile(userID, request.Nickname, request.Birth)
}

func (s *UserService) UpdateProfile(userID uint, data map[string]interface{}) error {
	var validData map[string]interface{}

	if nickname, ok := data["nickname"]; ok {
		err := s.repo.CheckNickname(nickname.(string))
		if err != nil {
			return err
		}
	}

	profile := &models.Profile{}
	profileData := profile.ConvToMap()

	for k, v := range data {
		if _, ok := profileData[k]; ok {
			validData[k] = v
		}
	}
	return s.repo.UpdateUserProfile(userID, validData)
}

func (s *UserService) UploadProfileImage(userID uint, file multipart.File, header *multipart.FileHeader) error {
	filepath := fmt.Sprintf("photos/%s", header.Filename)

	_, err := s.aws.UploadFile(file, header)
	if err != nil {
		return err
	}

	return s.repo.UploadUserProfileImage(userID, filepath)
}

func (s *UserService) DeleteUser(ID uint) error {
	return s.repo.DeleteUser(ID)
}

func (s *UserService) UpdatePassword(userID uint, newPassword string) error {
	return s.repo.UpdatePassword(userID, newPassword)
}

func (s *UserService) GetFollowings(userID uint) (models.Users, error) {
	users, err := s.repo.GetFollowings(userID)
	if err != nil {
		return nil, err
	}

	for idx, user := range users {
		users[idx].Profile.ProfileImage = s.aws.GetFilePath(user.Profile.ProfileImage)
	}

	return users, nil
}

func (s *UserService) GetFollowers(userID uint) (models.Users, error) {
	users, err := s.repo.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	for idx, user := range users {
		users[idx].Profile.ProfileImage = s.aws.GetFilePath(user.Profile.ProfileImage)
	}

	return users, nil
}

func (s *UserService) Follow(userID, followerID uint) error {
	return s.repo.Follow(userID, followerID)
}

func (s *UserService) UnFollow(userID, followingID uint) error {
	return s.repo.UnFollow(userID, followingID)
}

func (s *UserService) GetFavoriteQuants(userID uint) ([]*models.Quant, error) {
	return s.repo.GetFavoriteQuants(userID)
}

func (s *UserService) AddToFavoriteQuants(userID, quantID uint) error {
	return s.repo.AddToFavoriteQuants(userID, quantID)
}

func (s *UserService) DeleteFromFavoriteQuants(userID, quantID uint) error {
	return s.repo.DeleteFromFavoriteQuants(userID, quantID)
}
