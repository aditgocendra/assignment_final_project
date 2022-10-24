package middleware

import (
	"final_project/database"
	"final_project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizationData(param, tableName string) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		db := database.GetDB()
		paramId, err := strconv.Atoi(ctx.Param(param))
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Request",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		
		var modelUserID uint

		switch tableName {
			case "photos":
				Photo := models.Photo{}
				err = db.Select("user_id").First(&Photo, uint(paramId)).Error
				modelUserID = Photo.UserID
			case "social_media":
				SocialMedia := models.SocialMedia{}
				err = db.Select("user_id").First(&SocialMedia, uint(paramId)).Error
				modelUserID = SocialMedia.UserID
			case "comments":
				Comment := models.Comment{}
				err = db.Select("user_id").First(&Comment, uint(paramId)).Error
				modelUserID = Comment.UserID
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data not found",
				"message" : "data doesn't exist",
			})
			return
		}

		
		if modelUserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorization",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}