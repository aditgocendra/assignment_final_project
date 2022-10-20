package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social media is required, url~Url format not valid"`
	UserID         uint
	User		   *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	// Check all field create
	_, errValidate := govalidator.ValidateStruct(s)

	if errValidate != nil {
		err = errValidate
		return
	}

	return nil
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check all field create
	_, errValidate := govalidator.ValidateStruct(s)

	if errValidate != nil {
		err = errValidate
		return
	}

	return nil
}