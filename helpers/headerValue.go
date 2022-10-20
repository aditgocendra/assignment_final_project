package helpers

import "github.com/gin-gonic/gin"

var  (
	AppJSON = "application/json"
)

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}