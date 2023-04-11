package middleware

import (
	"gin-goinc-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddlelware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message" : "token is required",
		})
		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message" : "token is required",
		})
		return
	}

	claimsData, err := utils.DecodeToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message" : "token is invalid",
		})
	}

	ctx.Set("claimsData", claimsData)
	ctx.Set("user_id", claimsData["id"])
	ctx.Set("user_name", claimsData["name"])
	ctx.Set("user_email", claimsData["email"])


	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message" : "token is required",
		})
		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message" : "token is invalid",
		})
	}

	ctx.Next()
}