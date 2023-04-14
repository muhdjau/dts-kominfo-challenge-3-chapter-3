package models

import (
	"challenge-chapter-3-sesi-3/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Users struct {
	GormModel
	FullName string     `json:"full_name" gorm:"not null" form:"full_name" valid:"required~Your full name is required"`
	Email    string     `json:"email" gorm:"not null;uniqueIndex" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string     `json:"password" gorm:"not null" form:"password" valid:"required~Your password is required, minstringlength(6)~Password minimum length of 6 characters"`
	Products []Products `json:"products" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	RoleID   uint       `json:"role_id" gorm:"not null" form:"role_id" valid:"required~Role is required"`
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}
