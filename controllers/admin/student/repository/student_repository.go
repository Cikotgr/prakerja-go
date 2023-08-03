package repository

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/student"
	"github.com/ardin2001/backend-pemilu/models"
	"gorm.io/gorm"
)

type StudentRepositoryInterface interface {
	GetAll(nim string) (*[]student.StudentResponse, error)
	GetById(id string) (*student.StudentResponse, error)
	Create(student *models.Student) error
	Update(student *models.Student) error
	Delete(id string) error
}

type StudentRepositoryStruct struct {
	DB *gorm.DB
}

func NewStudentRepository(DB *gorm.DB) StudentRepositoryInterface {
	return &StudentRepositoryStruct{
		DB: DB,
	}
}

func (sr *StudentRepositoryStruct) GetAll(nim string) (*[]student.StudentResponse, error) {
	var students []student.StudentResponse

	err := sr.DB.Table("students").Where("nim LIKE ?", "%"+nim+"%").Find(&students).Error
	if err != nil {
		return nil, err
	}
	return &students, nil
}

func (sr *StudentRepositoryStruct) GetById(id string) (*student.StudentResponse, error) {
	var student student.StudentResponse
	err := sr.DB.Table("students").Where("id =?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (sr *StudentRepositoryStruct) Create(student *models.Student) error {
	err := sr.DB.Create(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (sr *StudentRepositoryStruct) Update(student *models.Student) error {
	err := sr.DB.Updates(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (sr *StudentRepositoryStruct) Delete(id string) error {
	err := sr.DB.Transaction(func(tx *gorm.DB) error {
		// err := tx.Model(&entity.ReadingListArticle{}).Unscoped().Delete(&entity.ReadingListArticle{}, "reading_list_id = ?", id).Error

		// if err != nil {
		// 	return err
		// }
		err := tx.Model(&models.Student{}).Unscoped().Delete(&models.Student{}, "id = ?", id).Error

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
