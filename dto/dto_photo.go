package dto

import "time"

type PhotoReq struct {
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
}

type CreatePhotoRes struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	CreatedAt *time.Time
}

type UpdatePhotoRes struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	UpdatedAt *time.Time
}


type GetPhotoRes struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	User PhotoUserRes
}

type PhotoUserRes struct{
	Email string
	Username string
}

