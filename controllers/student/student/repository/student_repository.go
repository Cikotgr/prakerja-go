package repository

import (
	"github.com/ardin2001/backend-pemilu/controllers/student/student"
	"gorm.io/gorm"
)

type StudentRepositoryInterface interface {
	LoginStudent(nim string) (*student.StudentResponse, error)
}

type StudentRepositoryStruct struct {
	DB *gorm.DB
}

func NewStudentRepository(DB *gorm.DB) StudentRepositoryInterface {
	return &StudentRepositoryStruct{
		DB: DB,
	}
}

func (ar *StudentRepositoryStruct) LoginStudent(nim string) (*student.StudentResponse, error) {
	var student student.StudentResponse
	err := ar.DB.Table("students").Where("nim = ? ", nim).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
