package app_config

import "os"

var PORT = ":8000"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")

	if portEnv != ""{
		PORT = portEnv
	}
}