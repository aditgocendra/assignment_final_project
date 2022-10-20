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

// CreateSocialMedia godoc
// @Summary     Create Social Media
// @Description Note : To create social media, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Param       request body     dto.SocialMediaReq true "create social media"
// @Success     201     {object} dto.CreateSocialMediaRes
// @Failure     400     {object} dto.ErrorMessage
// @Failure     401     {object} dto.ErrorMessage
// @Router      /socialmedias/ [post]
// @Security    Bearer <JWT>
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userIdData := uint(userData["id"].(float64))

	CreateSocialMediaReq := dto.SocialMediaReq{}

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&CreateSocialMediaReq)
	} else {
		ctx.ShouldBind(&CreateSocialMediaReq)
	}

	SocialMedia := models.SocialMedia{
		Name: CreateSocialMediaReq.Name,
		SocialMediaUrl: CreateSocialMediaReq.SocialMediaUrl,
	}

	SocialMedia.UserID = userIdData
	
	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	CreateSocialMediaRes := dto.CreateSocialMediaRes{
		ID: int(SocialMedia.ID),
		Name: SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
		UserID: int(SocialMedia.UserID),
		CreatedAt: SocialMedia.CreatedAt,
	}

	ctx.JSON(http.StatusOK, &CreateSocialMediaRes)
}


// GetSocialMedia godoc
// @Summary     Get Social Media
// @Description Note : To get social media, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Success     201 {object} dto.SocialMedias
// @Failure     400 {object} dto.ErrorMessage
// @Failure     401 {object} dto.ErrorMessage
// @Router      /socialmedias/ [get]
// @Security    Bearer <JWT>
func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	SocialMedias := []models.SocialMedia{}

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username")
	  }).Find(&SocialMedias).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"social_medias" : SocialMedias,
	})
}

// UpdateSocialMedia godoc
// @Summary     Update Social Media
// @Description Note : To update social media, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Param       socialMediaId path     string             true "socialMediaId"
// @Param       request       body     dto.SocialMediaReq true "update social media"
// @Success     200           {object} dto.UpdateSocialMediaRes
// @Failure     400           {object} dto.ErrorMessage
// @Failure     401           {object} dto.ErrorMessage
// @Failure     404           {object} dto.ErrorMessage
// @Router      /socialmedias/{socialMediaId} [put]
// @Security    Bearer <JWT>
func UpdateSocialMedia(ctx *gin.Context)  {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	// Check convert id string to int
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : err.Error(),
		})
		return
	}	

	UpdateSocialMediaReq := dto.SocialMediaReq{}

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&UpdateSocialMediaReq)
	}else{
		ctx.ShouldBind(&UpdateSocialMediaReq)
	}

	SocialMedia := models.SocialMedia{
		Name: UpdateSocialMediaReq.Name,
		SocialMediaUrl: UpdateSocialMediaReq.SocialMediaUrl,
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err = db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(&SocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" : "Bad Request",
			"message" : err.Error(),	
		})
		return
	}

	UpdateSocialMediaRes := dto.UpdateSocialMediaRes{
		ID: int(SocialMedia.ID),
		Name: SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
		UserID: int(SocialMedia.UserID),
		UpdatedAt: SocialMedia.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, UpdateSocialMediaRes)

}

// DeleteSocialMedia godoc
// @Summary     Delete Social Media
// @Description Note : To delete social media, you need to send the token that has been obtained from the login response with the format Bearer <YOUR TOKEN>.
// @Tags        social_media
// @Accept      json
// @Produce     json
// @Param       socialMediaId path     string true "socialMediaId"
// @Success     200           {object} dto.SuccessMessage
// @Failure     400           {object} dto.ErrorMessage
// @Failure     401           {object} dto.ErrorMessage
// @Failure     404           {object} dto.ErrorMessage
// @Router      /socialmedias/{socialMediaId} [delete]
// @Security    Bearer <JWT>
func DeleteSocialMedia(ctx *gin.Context)  {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	SocialMedia := models.SocialMedia{}

	userID := uint(userData["id"].(float64))
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if contentType == helpers.AppJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	}else{
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	// Delete To Database
	deletedComment := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(&SocialMedia)
	
	// Check rows affected
	if deletedComment.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err" : "Not Found",
			"message" : "Data not found",
		})
		return
	}

	// Response Success
	SuccesRes := dto.SuccessMessage{
		Message: "Your social media has been successfully deleted",
	}
	ctx.JSON(http.StatusOK, SuccesRes)

}