package boostrap

import (
	"gin-goinc-api/configs/app_config"
	"gin-goinc-api/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	app := gin.Default()

	routes.InitRoute(app)
	
	app.Run(app_config.PORT)
}