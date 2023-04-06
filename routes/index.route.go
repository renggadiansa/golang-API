package routes

import (
	"gin-goinc-api/configs/app_config"
	"gin-goinc-api/controllers/book_controller"
	filecontroller "gin-goinc-api/controllers/book_controller/file_controller"
	"gin-goinc-api/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	//route user
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/paginate", user_controller.GetUserPaginate)
	route.POST("/user", user_controller.Store)
	route.GET("/user/:id", user_controller.GetById)
	route.PATCH("/user/:id", user_controller.UpdateById)
	route.DELETE("/user/:id", user_controller.DeleteById)


	//route book
	route.GET("/book", book_controller.GetAllBook)

	//route controller file
	route.POST("/file", filecontroller.HandleUploadFile)
	route.DELETE("/file/:filename", filecontroller.HandleRemoveFile)

}