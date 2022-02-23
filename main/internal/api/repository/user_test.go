package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestUserRepo_GetUserByID_With_Valid_ID tests GetUserDataWithFields with valid data
func TestUserRepo_GetUserByID_With_Valid_ID(t *testing.T) {
	var id uint
	a := assert.New(t)
	for id = 1; id <= NumberOfTestCase; id++ {
		user, err := userRepo.GetUserByID(id)
		a.NoError(err, fmt.Sprintf("error generated while testing GetUserDataWithFields of id '%d'", id))
		a.NotEqual(uint(0), user.Profile.UserID, "profile.user_id must not be 0")
	}
}

func TestUserRepo_GetUserByID_With_Invalid_ID(t *testing.T) {
	var id uint
	a := assert.New(t)
	for id = NumberOfTestCase * 2; id <= NumberOfTestCase*3; id++ {
		_, err := userRepo.GetUserByID(id)
		a.Error(err, fmt.Sprintf("error should be generated while testing GetUserDataWithFields of id '%d'", id))
	}
}

func TestUserRepo_CreateUserAndProfile_With_Valid_InputData(t *testing.T) {
	a := assert.New(t)
	validData := map[string]interface{}{
		"email":    "test-valid@gmail.com",
		"password": "1234",
		"name":     "test-valid",
		"nickname": "test-valid",
		"birth":    "1997-11-10",
	}
	err := userRepo.CreateUserAndProfile(validData)
	a.NoError(err, "error should not be generated: valid data")
}

func TestUserRepo_CreateUserAndProfile_With_Invalid_InputData_Empty_Birth(t *testing.T) {
	a := assert.New(t)
	invalidData := map[string]interface{}{
		"email":    "test-invalid@gmail.com",
		"password": "1234",
		"name":     "test-invalid-01",
		"nickname": "test-invalid-01",
	}
	err := userRepo.CreateUserAndProfile(invalidData)
	a.Error(err, fmt.Sprintf("error must be generated: missing birth"))
}

func TestUserRepo_CreateUserAndProfile_With_Invalid_InputData_Duplicated_Email(t *testing.T) {
	a := assert.New(t)
	invalidData := map[string]interface{}{
		"email":    "jiheo@student.42seoul.kr",
		"password": "1234",
		"name":     "test-invalid",
		"nickname": "test-invalid",
		"birth":    "1997-11-10",
	}
	err := userRepo.CreateUserAndProfile(invalidData)
	a.Error(err, fmt.Sprintf("error must be generated: duplicated email"))
}

func TestUserRepo_CreateUserAndProfile_With_Invalid_InputData_Duplicated_Nickname(t *testing.T) {
	a := assert.New(t)
	invalidData := map[string]interface{}{
		"email":    "test-invalid@gmail.com",
		"password": "1234",
		"name":     "test-invalid",
		"nickname": "ex00",
		"birth":    "1997-11-10",
	}
	err := userRepo.CreateUserAndProfile(invalidData)
	a.Error(err, fmt.Sprintf("error must be generated: duplicated nickname"))
}

func TestUserRepo_UpdateUserProfile_With_Valid_Data(t *testing.T) {
	a := assert.New(t)
	parsedTime, _ := time.Parse("2006-01-02", "2000-01-01")
	profile := map[string]interface{}{
		"nickname":         "edited",
		"birth":            parsedTime,
		"email":            "contact@gmail.com",
		"phone":            "01001010101",
		"user_url":         "https://github.com/Jinseok-Heo",
		"intro_message":    "Hello world!(edited)",
		"location_city":    "Seoul",
		"location_country": "Korea",
	}
	err := userRepo.UpdateUserProfile(1, profile)
	a.NoError(err, fmt.Errorf("error must not be generated - valid data:%w", err))
}

func TestUserRepo_UpdateUserProfile_With_Invalid_Data_Duplicated_Nickname(t *testing.T) {
	a := assert.New(t)
	profile := map[string]interface{}{
		"nickname": "ex00",
	}
	err := userRepo.UpdateUserProfile(1, profile)
	a.Error(err, "error must be generated: duplicated nickname")
}

func TestUserRepo_DeleteUser(t *testing.T) {
	a := assert.New(t)
	err := userRepo.DeleteUser(10)
	a.NoError(err, "error while deleting a user with id 10")

	_, err = userRepo.GetUserByID(10)
	a.Error(err, "error while getting a deleted user with id 10")
}

func TestUserRepo_Follow(t *testing.T) {
	a := assert.New(t)
	err := userRepo.Follow(1, 2)
	a.NoError(err, "error while following user 2")
	err = userRepo.Follow(1, 3)
	a.NoError(err, "error while following user 3")
	err = userRepo.Follow(1, 5)
	a.NoError(err, "error while following user 5")
	err = userRepo.Follow(1, 7)
	a.NoError(err, "error while following user 7")
}

func TestUserRepo_Follow_Invalid_User(t *testing.T) {
	a := assert.New(t)
	err := userRepo.Follow(1, 30)
	a.Error(err, "error should be generated while following not existing user id=30")
}

func TestUserRepo_GetFollowers(t *testing.T) {
	a := assert.New(t)
	_ = userRepo.Follow(3, 5)
	_ = userRepo.Follow(3, 7)
	_ = userRepo.Follow(3, 6)
	followers, err := userRepo.GetFollowers(3)
	a.NoError(err, "error should not be generated while getting followers of user 3")
	a.NotEmpty(followers, "followers should not be empty")
	if followers != nil && len(followers) != 0 {
		a.Equal(uint(1), followers[0].ID, "should be equal with 1 and follower's id")
		a.NotEqual(uint(0), followers[0].Profile.UserID, "profile.user_id must not be 0")
	}
}

func TestUserRepo_GetFollowings(t *testing.T) {
	a := assert.New(t)
	followings, err := userRepo.GetFollowings(1)
	a.NoError(err, fmt.Errorf("error should not be generated while getting followings of user 1: %w", err))
	a.Equal(4, len(followings))
}

func TestUserRepo_UnFollow(t *testing.T) {
	a := assert.New(t)
	err := userRepo.UnFollow(1, 2)
	a.NoError(err, fmt.Errorf("error should not be generated while unfollowing user 2: %w", err))
}

func TestUserRepo_GetFavoriteQuants(t *testing.T) {
	a := assert.New(t)
	_, err := userRepo.GetFavoriteQuants(1)
	a.NoError(err, err)
}

func TestUserRepo_AddToFavoriteQuants(t *testing.T) {
	a := assert.New(t)
	err := userRepo.AddToFavoriteQuants(1, 2)
	a.NoError(err, err)
	err = userRepo.AddToFavoriteQuants(1, 3)
	a.NoError(err, err)
	err = userRepo.AddToFavoriteQuants(1, 5)
	a.NoError(err, err)
	err = userRepo.AddToFavoriteQuants(1, 7)
	a.NoError(err, err)
}

func TestUserRepo_DeleteFromFavoriteQuants(t *testing.T) {
	a := assert.New(t)
	err := userRepo.DeleteFromFavoriteQuants(1, 5)
	a.NoError(err, err)
}
