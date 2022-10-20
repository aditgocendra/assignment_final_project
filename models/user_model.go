package models

import (
	"errors"
	"final_project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username 	string 			`gorm:"not null" json:"username,omitempty" form:"username" valid:"required~Your username is required"`
	Email    	string 			`gorm:"not null;uniqueIndex" json:"email,omitempty" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password 	string			`gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 character"`
	Age      	int				`gorm:"not null" json:"age,omitempty" form:"age" valid:"required~Your age is required, range(8|100)~Age minimum 8"`
	Photo 		[]Photo			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo,omitempty" valid:"-"`
	Comment 	[]Comment 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment,omitempty" valid:"-"`
	SocialMedia []SocialMedia 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media,omitempty" valid:"-"`
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error)  {
	// Check all field create
	_, errValidate := govalidator.ValidateStruct(u)
	
	if errValidate != nil {
		err = errValidate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if govalidator.IsNull(u.Username) {
		err = errors.New("Username is required")
		return
	}

	if govalidator.IsNull(u.Email) {
		err = errors.New("Email is required")
		return
	}

	if !govalidator.IsEmail(u.Email) {
		err = errors.New("Invalid email format")
		return
	}

	return nil


}

