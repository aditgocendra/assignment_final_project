package controllers

import (
	"final_project/database"
	"final_project/dto"
	"final_project/helpers"
	"final_project/models"
	"strconv"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary     Register User
// @Description User registration, your password is stored securely because it has gone through the encryption process
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     dto.UserCreateReq true "create user"
// @Success     201     {object} dto.UserCreateRes
// @Failure     400     {object} dto.ErrorMessage
// @Router      /users/register [post]
func Register(ctx *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(ctx)

	UserCreateReq := dto.UserCreateReq{}

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&UserCreateReq)
	}else{
		ctx.ShouldBind(&UserCreateReq)
	}

	UserModel := models.User{
		Username: UserCreateReq.Username,
		Email: UserCreateReq.Email,
		Password: UserCreateReq.Password,
		Age: UserCreateReq.Age,
	}
	
	err := db.Debug().Create(&UserModel).Error

	if err != nil {
		errMessage := dto.ErrorMessage{
			TypeError: "Bad Request",
			Message:  "Invalid format request",
		}
		ctx.JSON(http.StatusBadRequest, errMessage)
		return
	}

	// Response Success Create
	UserCreateRes := dto.UserCreateRes{
		ID: int(UserModel.ID),
		Email: UserModel.Email,
		Username : UserModel.Username,
		Age: UserModel.Age,
	}

	ctx.JSON(http.StatusCreated, UserCreateRes)	
}

// LoginUser godoc
// @Summary     Login User
// @Description Login your account after register
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body     dto.UserLoginReq true "create user"
// @Success     201     {object} dto.UserLoginRes
// @Failure     401     {object} dto.ErrorMessage
// @Failure     400     {object} dto.ErrorMessage
// @Router      /users/login [post]
func Login(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	UserRequest := dto.UserLoginReq{}

	// Convert data to struct user
	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&UserRequest)
	}else{
		ctx.ShouldBind(&UserRequest)
	}

	User := models.User{
		Email: UserRequest.Email,
		Password: UserRequest.Password,
	}

	password := User.Password

	// Get data user login from db
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized",
			"message" : "Invalid email",
		})
		return
	}

	// Compare password hash
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized",
			"message" : "Invalid password",
		})
		return
	}
	
	// Generate token
	token := helpers.GenerateToken(User.ID, User.Email)
	
	LoginResponse := dto.UserLoginRes{
		Token: token,
	}

	// Success
	ctx.JSON(http.StatusOK, LoginResponse)

}

// UpdateUser godoc
// @Summary     Update User
// @Description Note : To change your account data, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       id      path     string            true "id"
// @Param       request body     dto.UserUpdateReq true "create user"
// @Success     200     {object} dto.UserUpdateRes
// @Failure     400     {object} dto.ErrorMessage
// @Failure     401     {object} dto.ErrorMessage
// @Failure     403     {object} dto.ErrorMessage
// @Failure     404     {object} dto.ErrorMessage
// @Router      /users/{id} [put]
// @Security    Bearer <JWT>
func UpdateUser(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	userIdParam, err := strconv.Atoi(ctx.Param("id"))
	
	// Check convert id string to int
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : "Request fail, param type not valid",
		})
		return
	}	
	UserUpdateReq := dto.UserUpdateReq{}

	// Convert data to struct user
	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&UserUpdateReq)
	}else{
		ctx.ShouldBind(&UserUpdateReq)
	}

	userIdData := uint(userData["id"].(float64))

	// Check user data id JWT and param user id
	if userIdData != uint(userIdParam) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"err" : "Forbidden",
			"message" : "You are not the owner of this account",
		})
		return
	}

	UserModel := models.User{
		Username:  UserUpdateReq.Username,
		Email:  UserUpdateReq.Email,

	}

	// Update To Database
	err = db.Model(&UserModel).Where("id = ?", userIdParam).Updates(&UserModel).Scan(&UserModel).Error
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : "Invalid format request",
		})
		return
	}

	// Success Response
	UpdateUserRes := dto.UserUpdateRes{
		ID: UserModel.ID,
		Email: UserModel.Email,
		Username: UserModel.Username,
		Age: UserModel.Age,
		UpdatedAt: UserModel.UpdatedAt,
	}
	
	ctx.JSON(http.StatusOK, UpdateUserRes)
}

// DeleteUser godoc
// @Summary     Delete User
// @Description Note : To delete your account data, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        user
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.SuccessMessage
// @Failure     400 {object} dto.ErrorMessage
// @Failure     401 {object} dto.ErrorMessage
// @Failure     403 {object} dto.ErrorMessage
// @Failure     404 {object} dto.ErrorMessage
// @Router      /users/ [delete]
// @Security    Bearer <JWT>
func DeleteUser(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	UserModel := models.User{}

	// Convert data to struct user
	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&UserModel)
	}else{
		ctx.ShouldBind(&UserModel)
	}

	userIdData := uint(userData["id"].(float64))
	UserModel.ID = userIdData
	
	// Delete To Database
	deletedUser := db.Model(&UserModel).Delete(&UserModel)
	
	// Check rows affected
	if deletedUser.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err" : "Not Found",
			"message" : "Data not found",
		})
		return
	}

	// Check error delete data user
	if deletedUser.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : "Request delete fail",
		})
		return
	}
	
	// Success Deleted
	UserDeleted := dto.SuccessMessage{
		Message: "Your account has been successfully deleted",
	}
	ctx.JSON(http.StatusOK, UserDeleted)
}