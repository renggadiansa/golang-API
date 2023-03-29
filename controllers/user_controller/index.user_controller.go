package user_controller

import (
	"gin-goinc-api/database"
	"gin-goinc-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
						"message": "Internal Server Error",
					})
					return
	}
	
	// isValidated := true

	// 	if !isValidated {
	// 		ctx.AbortWithStatusJSON(400, gin.H{
	// 			"message": "Bad Request",
	// 		})
	// 		return
	// 	}

		ctx.JSON(200, gin.H{
			"data": users,
		})
}

func GetById(ctx *gin.Context)  {
			id := ctx.Param("id")
			user := new(models.User)
			
			errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
			if errDb != nil || user.ID == nil{
				ctx.JSON(http.StatusNotFound, gin.H{
					"messange" : "Data user not found",
				})
				return
			}

			ctx.JSON(200, gin.H {
				"message" : "data transmitted",
				"data" : user,
			})
}

func Store(ctx *gin.Context)  {
			
}

func UpdateById(ctx *gin.Context)  {
			
}

func DeleteById(ctx *gin.Context)  {
			
}