package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConfigDatabase() (*gorm.DB, error) {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST_LOCAL")
	dbUsername := os.Getenv("DB_USERNAME_LOCAL")
	dbPassword := os.Getenv("DB_PASSWORD_LOCAL")
	dbName := os.Getenv("DB_NAME_LOCAL")
	dbPort := os.Getenv("DB_PORT_LOCAL")
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
