package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint
	User 	*User
	PhotoID uint   `json:"photo_id" form:"photo_id"`
	Photo 	*Photo
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	// Check all field create
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	return nil
}


func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check all field create
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	return nil
}