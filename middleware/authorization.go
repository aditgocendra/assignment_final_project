package middleware

import (
	"final_project/database"
	"final_project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(ctx.Param("photoid"))
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Request",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data not found",
				"message" : "data doesn't exist",
			})
			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorization",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(ctx.Param("commentid"))
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : err.Error(),
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data not found",
				"message" : "data doesn't exist",
			})
			return
		}

		if Comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorization",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(ctx.Param("socialMediaId"))
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Request",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(commentId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data not found",
				"message" : "data doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorization",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}