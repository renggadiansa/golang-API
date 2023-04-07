package middleware

import "github.com/gin-gonic/gin"

func AuthMiddlelware(ctx *gin.Context) {
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