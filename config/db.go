package config

import (
	"log"
	"os"
	"strings"

	"sim-klinikv2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *ConfigDB) {
	str := []string{
		config.DBUserName,
		config.DBUserPassword,
		"@tcp(",
		config.DBHost,
		":",
		config.DBPort,
		")/",
		config.DBName,
		"?charset=utf8mb4&parseTime=True&loc=Local",
	}

	var err error
	dsn := strings.Join(str, "")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	// DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.AutoMigrate(&models.User{})

	log.Println("🚀 Connected Successfully to the Database")
}
