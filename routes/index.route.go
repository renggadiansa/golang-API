package routes

import (
	"gin-goinc-api/controllers/book_controller"
	"gin-goinc-api/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.GET("/user", user_controller.GetAllUser)
	route.GET("/book", book_controller.GetAllBook)
}