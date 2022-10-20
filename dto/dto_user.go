package dto

import "time"

type UserCreateReq struct {
	Email    string `json:"email" form:"email"`
	Username string	`json:"username" form:"username"`
	Password string	`json:"password" form:"password"`
	Age      int	`json:"age" form:"age"`
}

type UserCreateRes struct {
	ID		 int
	Email    string 
	Username string	
	Age      int	
}

type UserLoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string	`json:"password" form:"password"`
}

type UserLoginRes struct {
	Token string
}

type UserUpdateReq struct {
	Email     string `json:"email" form:"email"`	
	Username  string `json:"username" form:"username"`
}

type UserUpdateRes struct {
	ID		  uint
	Email     string `json:"email" form:"email"`
	Username  string `json:"username" form:"username"`
	Age       int	 `json:"age" form:"age"`
	UpdatedAt *time.Time 
}

