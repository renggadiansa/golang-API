package routes

import (
	"gin-goinc-api/middleware"

	"github.com/gin-gonic/gin"
	
	filecontroller "gin-goinc-api/controllers/book_controller/file_controller"
)

func v1Route(app *gin.RouterGroup) {
	route := app

	//route controller file
	authRoute := route.Group("file", middleware.AuthMiddlelware)
	authRoute.POST("/file",filecontroller.HandleUploadFile)
	authRoute.POST("/file/middleware", middleware.UploadFile ,filecontroller.SendStatus)
	authRoute.DELETE("/:filename",filecontroller.HandleRemoveFile)
}