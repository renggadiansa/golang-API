package configs

import (
	"gin-goinc-api/configs/app_config"
	"gin-goinc-api/configs/db_config"
)

func InitConfig() {

	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()

}