package database

import (
	"fmt"
	"gin-goinc-api/configs/db_config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDatabase() {

	var errConnection error

	if db_config.DB_DRIVER == "mysql" {

		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER,db_config.DB_PASSWORD,db_config.DB_HOST,db_config.DB_PORT,db_config.DB_NAME)

		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if db_config.DB_DRIVER == "pgsql" {

		dsnMysql := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", db_config.DB_HOST,db_config.DB_USER,db_config.DB_PASSWORD,db_config.DB_NAME,db_config.DB_PORT)

		DB, errConnection = gorm.Open(postgres.Open(dsnMysql), &gorm.Config{})
	}

	if errConnection != nil {
		panic("Failed to connect to database")
	} 

	log.Println("Connect to database")
	
}