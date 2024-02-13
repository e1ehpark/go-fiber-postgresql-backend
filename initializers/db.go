package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/e1ehpark/go-fiber-postgresql-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(1)
	}
	
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	DB.Logger = logger.Default.Logmode(logger.Info)

	log.Println("Running migrations...")
	DB.AutoMigrate(&models.Note{})

	log.Println("Connected Successfully to the Database")
}