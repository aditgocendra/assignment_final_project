package dto

import "time"

type CommentReq struct {
	Message string `json:"message" form:"message"`
	PhotoID int    `json:"photo_id" form:"photo_id"`
}

type CommentRes struct {
	ID        int
	Message   string
	PhotoID   int
	UserID    int
	CreatedAt *time.Time
}

type UpdateCommentReq struct {
	Message string `json:"message" form:"message"`
}

type GetCommentRes struct {
	ID        int
	Message   string
	PhotoID   int
	UserID    int
	CreatedAt *time.Time
	UpdatedAt  *time.Time
	User CommentUserRes
	Photo CommentPhotoRes
}

type CommentUserRes struct {
	ID 		 int
	Email 	 string
	Username string
}

type CommentPhotoRes struct {
	ID 			int
	Title 		string
	Caption 	string
	PhotoUrl 	string
	UserID 		int
}