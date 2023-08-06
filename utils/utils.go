package utils

import (
	"fmt"

	config "github.com/ardin2001/backend-pemilu/configs"
	"github.com/ardin2001/backend-pemilu/models"
)

func MigrateDB() {
	DB, err := config.ConfigDatabase()
	if err != nil {
		fmt.Println("Failed connect to database : ", err.Error())
		return
	}

	DB.AutoMigrate(&models.Role{}, &models.Student{}, &models.Admin{}, &models.Candidate{}, &models.CandidateStudent{})
}
