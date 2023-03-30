package user_controller

import (
	"gin-goinc-api/database"
	"gin-goinc-api/models"
	"gin-goinc-api/requests"
	"gin-goinc-api/responses"
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
			user := new(responses.UserResponse)
			
			errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

			if errDb != nil || user.ID == nil{
				ctx.JSON(500, gin.H{
					"messange" : "Inernal server error",
				})
				return
			}

			if user.ID == nil{
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
			userReq := new(requests.UserRequest)

			if errReq := ctx.ShouldBind(&userReq); errReq != nil {
				ctx.JSON(400, gin.H {
					"message" : errReq.Error(),
				})
				return
			}

			user := new(models.User)
			user.Name = &userReq.Name
			user.Address = &userReq.Address
			user.BornDate = &userReq.BornDate

			ErrDB := database.DB.Table("users").Create(&user).Error
			if ErrDB != nil {
				ctx.JSON(500, gin.H {
					"message" : "can't create data",
				})
				return
			}


			
			ctx.JSON(200, gin.H {
					"message" : "Data created successfully",
					"data" : user,
				})
}

func UpdateById(ctx *gin.Context)  {
			
}

func DeleteById(ctx *gin.Context)  {
			
}