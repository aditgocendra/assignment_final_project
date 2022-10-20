package dto

import "time"

type SocialMediaReq struct {
	Name           string `json:"name" form:"name"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url"`
}

type CreateSocialMediaRes struct {
	ID			   int
	Name           string
	SocialMediaUrl string
	UserID         int
	CreatedAt      *time.Time
}

type UpdateSocialMediaRes struct {
	ID			   int
	Name           string
	SocialMediaUrl string
	UserID         int
	UpdatedAt      *time.Time
}

type SocialMedias struct {
	ID			   int
	Name           string
	SocialMediaUrl string
	UserID         int
	CreatedAt      *time.Time
	UpdatedAt	   *time.Time
	User 		   SocialMediaUserRes
}

type SocialMediaUserRes struct {
	ID int
	Username string
}