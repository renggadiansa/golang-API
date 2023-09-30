package cors_config

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var origins = []string{
	"http://localhost:8080",
}

func CorsConfig(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credential", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

func CorsConfigContrib() gin.HandlerFunc {
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	config.AllowOrigins = origins

	return cors.New(config)

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Credential", "true")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
	// c.Writer.Header().Set("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

	// if c.Request.Method == "OPTIONS" {
	// 	c.AbortWithStatus(http.StatusNoContent)
	// 	return
	// }

	// c.Next()
}
