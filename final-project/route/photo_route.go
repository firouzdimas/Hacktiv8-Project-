package route

import (
	"github.com/firouzdimas/Hacktiv8-Project-/controller"
	"github.com/firouzdimas/Hacktiv8-Project-/middleware"
	"github.com/firouzdimas/Hacktiv8-Project-/repository"
	"github.com/firouzdimas/Hacktiv8-Project-/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPhotoRoute(router *gin.Engine, db *gorm.DB) {
	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	photoService := service.NewPhotoService(photoRepository, commentRepository)
	photoController := controller.NewPhotoController(photoService)

	authUser := router.Group("/photos", middleware.AuthMiddleware)
	{
		authUser.POST("", photoController.CreatePhoto)
		authUser.GET("", photoController.GetAll)
		authUser.GET("/:photo_id", photoController.GetOne)
		authUser.PUT("/:photo_id", photoController.UpdatePhoto)
		authUser.DELETE("/:photo_id", photoController.DeletePhoto)
	}
}
