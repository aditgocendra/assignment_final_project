package controllers

import (
	"final_project/database"
	"final_project/dto"
	"final_project/helpers"
	"final_project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateComment godoc
// @Summary     Create Comment
// @Description Note : To create comment, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        comment
// @Accept      json
// @Produce     json
// @Param       request body     dto.CommentReq true "create comment"
// @Success     201     {object} dto.CommentRes
// @Failure     400     {object} dto.ErrorMessage
// @Failure     401     {object} dto.ErrorMessage
// @Router      /comments [post]
// @Security    Bearer <JWT>
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userIdData := uint(userData["id"].(float64))

	CommentReq := dto.CommentReq{}

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&CommentReq)
	} else {
		ctx.ShouldBind(&CommentReq)
	}

	Comment := models.Comment{
		Message: CommentReq.Message,
		PhotoID: uint(CommentReq.PhotoID),
	}

	Comment.UserID = userIdData
	
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Success Response
	CommentRes := dto.CommentRes{
		ID:        int(Comment.ID),
		Message:   Comment.Message,
		PhotoID:   int(Comment.PhotoID),
		UserID:    int(Comment.UserID),
		CreatedAt: Comment.CreatedAt,
	}
	ctx.JSON(http.StatusOK, CommentRes)
}


// GetComments godoc
// @Summary     Get Comments
// @Description Note : To get comment, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        comment
// @Accept      json
// @Produce     json
// @Success     201 {object} dto.GetCommentRes
// @Failure     400 {object} dto.ErrorMessage
// @Failure     401 {object} dto.ErrorMessage
// @Router      /comments/ [get]
// @Security    Bearer <JWT>
func GetComment(ctx *gin.Context) {
	db := database.GetDB()

	Comments := []models.Comment{}

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	  }).Preload("Photo", func (db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Title", "Caption", "PhotoUrl", "UserID")
	  }).Find(&Comments).Error
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	GetCommentResp := []dto.GetCommentRes{}

	for _, v := range Comments {
		GetCommentRes := dto.GetCommentRes{
			ID: int(v.ID),
			Message: v.Message,
			PhotoID: int(v.PhotoID),
			UserID: int(v.UserID),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		GetCommentRes.User.ID = int(v.User.ID)
		GetCommentRes.User.Email = v.User.Email
		GetCommentRes.User.Username = v.User.Username

		GetCommentRes.Photo.ID = int(v.Photo.ID)
		GetCommentRes.Photo.Title = v.Photo.Title
		GetCommentRes.Photo.Caption = v.Photo.Caption
		GetCommentRes.Photo.PhotoUrl = v.Photo.PhotoUrl
		GetCommentRes.Photo.UserID = int(v.Photo.UserID)

		GetCommentResp = append(GetCommentResp, GetCommentRes)
	}

	ctx.JSON(http.StatusCreated, &Comments)
}

// UpdateComment godoc
// @Summary     Update Comment
// @Description Note : To update comment, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        comment
// @Accept      json
// @Produce     json
// @Param       commentid path     string               true "commentid"
// @Param       request   body     dto.UpdateCommentReq true "udpate comment"
// @Success     201       {object} dto.UpdatePhotoRes
// @Failure     400       {object} dto.ErrorMessage
// @Failure     401       {object} dto.ErrorMessage
// @Failure     404       {object} dto.ErrorMessage
// @Router      /comments/{commentid} [put]
// @Security    Bearer <JWT>
func UpdateComment(ctx *gin.Context)  {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	commentId, err := strconv.Atoi(ctx.Param("commentid"))

	// Check convert id string to int
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : err.Error(),
		})
		return
	}	

	userID := uint(userData["id"].(float64))

	UpdateComment := dto.UpdateCommentReq{}

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&UpdateComment)
	}else{
		ctx.ShouldBind(&UpdateComment)
	}

	Comment := models.Comment{
		Message: UpdateComment.Message,
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	// Update Comment
	err = db.Debug().Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Scan(&Comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : err.Error(),	
		})
		return
	}

	Photo := models.Photo{}

	Photo.UserID = Comment.UserID
	Photo.ID = Comment.PhotoID

	// Get Photo
	err = db.Debug().Model(&Photo).Find(&Photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : err.Error(),	
		})
		return
	}

	// Success Response
	UpdatePhotoRes := dto.UpdatePhotoRes{
		ID: int(Photo.ID),
		Title: Photo.Title,
		Caption: Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
		UserID: int(Photo.UserID),
		UpdatedAt: Photo.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, UpdatePhotoRes)
}

// DeleteComment godoc
// @Summary     Delete Comment
// @Description Note : To delete comment, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        comment
// @Accept      json
// @Produce     json
// @Param       commentid path     string true "commentid"
// @Success     201       {object} dto.SuccessMessage
// @Failure     400       {object} dto.ErrorMessage
// @Failure     401       {object} dto.ErrorMessage
// @Failure     404       {object} dto.ErrorMessage
// @Router      /comments/{commentid} [delete]
// @Security    Bearer <JWT>
func DeleteComment(ctx *gin.Context)  {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))
	commentId, _ := strconv.Atoi(ctx.Param("commentid"))


	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&Comment)
	}else{
		ctx.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	// Delete To Database
	deletedComment := db.Model(&Comment).Where("id = ?", commentId).Delete(&Comment)
	
	// Check rows affected
	if deletedComment.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err" : "Not Found",
			"message" : "Data not found",
		})
		return
	}

	// Success Deleted
	CommentDeleted := dto.SuccessMessage{
		Message: "Your comment has been successfully deleted",
	}

	ctx.JSON(http.StatusOK, CommentDeleted)

}