package boostrap

import (
	"gin-goinc-api/configs"
	"gin-goinc-api/configs/app_config"
	"gin-goinc-api/configs/cors_config"
	"gin-goinc-api/database"
	"gin-goinc-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {

	//load env
	err:= godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//init config
	configs.InitConfig()

	//database connection
	database.ConnectDatabase()

	//init gin enggine
	app := gin.Default()

	//cors config
	app.Use(cors_config.CorsConfigContrib())
	// app.Use(cors_config.CorsConfig())


	//init routes
	routes.InitRoute(app)
	
	//run app
	app.Run(app_config.PORT)
}