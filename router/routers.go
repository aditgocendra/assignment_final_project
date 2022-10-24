package router

import (
	"final_project/controllers"
	"final_project/middleware"

	_ "final_project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("/users")
	{	
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
		userRouter.PUT("/:id", middleware.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/", middleware.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoid",  middleware.AuthorizationData("photoid", "photos") ,controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoid", middleware.AuthorizationData("photoid", "photos") ,controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.PUT("/:commentid", middleware.AuthorizationData("commentid", "comments"), controllers.UpdateComment)
		commentRouter.DELETE("/:commentid", middleware.AuthorizationData("commentid", "comments"), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.AuthorizationData("socialMediaId", "social_media"), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.AuthorizationData("socialMediaId", "social_media"), controllers.DeleteSocialMedia)
	}

	return r
}