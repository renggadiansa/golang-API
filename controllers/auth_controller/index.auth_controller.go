package auth_controller

import (
	"gin-goinc-api/database"
	"gin-goinc-api/models"
	"gin-goinc-api/requests"
	"gin-goinc-api/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {

	loginReq := new(requests.LoginRequest)

	if errReq := ctx.ShouldBind(&loginReq); errReq != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	user := new(models.User)
	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).First(&user).Error

	log.Println(errUser)

	if errUser != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	//check pass

	if loginReq.Password != "1234" {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "password not match",
		})
		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		ctx.AbortWithStatusJSON(500, gin.H{

			"message": "failed to generate token",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "login success",
		"token":   token,
	})
}
