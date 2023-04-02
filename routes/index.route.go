package routes

import (
	"gin-goinc-api/controllers/book_controller"
	"gin-goinc-api/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	//route user
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/paginate", user_controller.GetUserPaginate)
	route.POST("/user", user_controller.Store)
	route.GET("/user/:id", user_controller.GetById)
	route.PATCH("/user/:id", user_controller.UpdateById)
	route.DELETE("/user/:id", user_controller.DeleteById)


	route.GET("/book", book_controller.GetAllBook)
}