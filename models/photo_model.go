package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `gorm:"not null" json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo url is required, url~Url format not valid"`
	UserID   uint
	User	 *User 	`json:",omitempty"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	// Check all field create
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	return nil
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	// Check all field update
	_, errValidate := govalidator.ValidateStruct(p)

	if errValidate != nil {
		err = errValidate
		return
	}

	return nil
}