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

// CreatePhoto godoc
// @Summary     Create Photo
// @Description Note : To create photo, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        photo
// @Accept      json
// @Produce     json
// @Param       request body     dto.PhotoReq true "create user"
// @Success     201     {object} dto.CreatePhotoRes
// @Failure     400     {object} dto.ErrorMessage
// @Failure     401     {object} dto.ErrorMessage
// @Router      /photos/ [post]
// @Security    Bearer <JWT>
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB() 

	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userIdData := uint(userData["id"].(float64))

	PhotoReq := dto.PhotoReq{}


	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&PhotoReq)
	} else {
		ctx.ShouldBind(&PhotoReq)
	}

	Photo := models.Photo{
		Title: PhotoReq.Title,
		Caption: PhotoReq.Caption,
		PhotoUrl: PhotoReq.PhotoUrl,
	}

	Photo.UserID = userIdData
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Request create photo fail",
		})
		return
	}

	// Success Response
	PhotoRes := dto.CreatePhotoRes{
		ID: int(Photo.ID),
		Title: Photo.Title,
		Caption: Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
		UserID : int(Photo.UserID),
		CreatedAt: Photo.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, PhotoRes)
}

// GetPhotos godoc
// @Summary     Get Photo
// @Description Note : To get photos, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        photo
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.GetPhotoRes
// @Failure     400 {object} dto.ErrorMessage
// @Failure     401 {object} dto.ErrorMessage
// @Router      /photos/ [get]
// @Security    Bearer <JWT>
func GetPhotos(ctx *gin.Context) {
	db := database.GetDB()

	Photos := []models.Photo{}
	
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Order("Username DESC").Select("ID","Username", "Email")
	  }).Find(&Photos).Error
	
	GetPhotoRes := []dto.GetPhotoRes{}

	for _, v := range Photos {
		PhotoRes := dto.GetPhotoRes{
			ID : int(v.ID),
			Title : v.Title,
			Caption : v.Caption,
			PhotoUrl : v.PhotoUrl,
			UserID : int(v.UserID),
			CreatedAt : v.CreatedAt,
			UpdatedAt : v.UpdatedAt,
		}

		PhotoRes.User.Email = v.User.Email
		PhotoRes.User.Username = v.User.Username
		GetPhotoRes = append(GetPhotoRes, PhotoRes)
	}
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Request get photo fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, &GetPhotoRes)
}

// UpdatePhoto godoc
// @Summary     Update Photo
// @Description Note : To update photos, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        photo
// @Accept      json
// @Produce     json
// @Param       photoid path     string       true "photoid"
// @Param       request body     dto.PhotoReq true "update photo"
// @Success     201     {object} dto.UpdatePhotoRes
// @Failure     400     {object} dto.ErrorMessage
// @Failure     401     {object} dto.ErrorMessage
// @Failure     404     {object} dto.ErrorMessage
// @Router      /photos/{photoid} [put]
// @Security    Bearer <JWT>
func UpdatePhoto(ctx *gin.Context)  {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Photo := models.Photo{}

	photoId, err := strconv.Atoi(ctx.Param("photoid"))

	// Check convert id string to int
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message": "Request fail, param not valid",
		})
		return
	}	

	userID := uint(userData["id"].(float64))

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&Photo)
	}else{
		ctx.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : "Request update photo fail",	
		})
		return
	}

	ctx.JSON(http.StatusOK, Photo)

}

// DeletePhoto godoc
// @Summary     Delete Photo
// @Description Note : To delete photos, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        photo
// @Accept      json
// @Produce     json
// @Param       photoid path     string true "photoid"
// @Success     200     {object} dto.SuccessMessage
// @Failure     400     {object} dto.ErrorMessage
// @Failure     401     {object} dto.ErrorMessage
// @Failure     404     {object} dto.ErrorMessage
// @Router      /photos/{photoid} [delete]
// @Security    Bearer <JWT>
func DeletePhoto(ctx *gin.Context)  {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(ctx.Param("photoid"))
	userID := uint(userData["id"].(float64))
	
	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&Photo)
	}else{
		ctx.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	// Delete To Database
	deletedPhoto := db.Model(&Photo).Where("id = ?", photoId).Delete(&Photo)
	
	// Check rows affected
	if deletedPhoto.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err" : "Not Found",
			"message" : "Data not found",
		})
		return
	}

	// Check error delete data user
	if deletedPhoto.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : "Request delete photo fail",
		})
		return
	}

	// Success Deleted
	PhotoDeleted := dto.SuccessMessage{
		Message: "Your photo has been successfully deleted",
	}

	ctx.JSON(http.StatusOK, PhotoDeleted)
}