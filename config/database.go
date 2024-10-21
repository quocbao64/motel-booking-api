package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func ConnectDB() (db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbHost, _ := os.LookupEnv("DB_HOST")
	dbUser, _ := os.LookupEnv("DB_USER")
	dbPassword, _ := os.LookupEnv("DB_PASS")
	dbName, _ := os.LookupEnv("DB_NAME")
	dbPort, _ := os.LookupEnv("DB_PORT")

	dst := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dst), &gorm.Config{
		TranslateError:                           true,
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		fmt.Println("Failed to connect to the database")
		return
	}

	return
}

func CloseDB(db *gorm.DB) {
	dbSQL, _ := db.DB()
	err := dbSQL.Close()
	if err != nil {
		return
	}
}
