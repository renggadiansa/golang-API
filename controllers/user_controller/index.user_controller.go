package user_controller

import (
	"gin-goinc-api/database"
	"gin-goinc-api/models"

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
			
}

func Store(ctx *gin.Context)  {
			
}

func UpdateById(ctx *gin.Context)  {
			
}

func DeleteById(ctx *gin.Context)  {
			
}