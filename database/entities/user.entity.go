package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model      `json:"-"`
	Username        string `json:"username" binding:"required" gorm:"column:username"`
	Email           string `json:"email" binding:"required" gorm:"column:email;unique:not null"`
	Password        string `json:"password,omitempty" binding:"required" gorm:"column:password"`
	ConfirmPassword string `json:"confirmPassword,omitempty" binding:"required" gorm:"-"`
}

func (UserEntity) TableName() string {
	return "users"
}
