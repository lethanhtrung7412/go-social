package operation

import (
	"go_social/database"
	e "go_social/database/entities"
	"go_social/internal/common"
)

var DB = database.DB()

func FetchUser(users *[]e.UserEntity) error {
	allUsers := DB.Model(&e.UserEntity{}).Order("username asc").Find(
		&users)
	if allUsers.Error != nil {
		return allUsers.Error
	}

	return nil
}

func FindOne(email string) (*e.UserEntity, error) {
	user := &e.UserEntity{}
	if err := DB.Model(&e.UserEntity{}).Where("email = ?",
		email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateNewUser(user *e.UserEntity) (uint, error) {
	pass, err := common.HashPassword(user.Password)

	if err != nil {
		return 0, err
	}
	user.Password = string(pass)
	createUser := DB.Model(&e.UserEntity{}).Create(&user)
	err = createUser.Error

	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
