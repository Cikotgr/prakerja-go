package usecase

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/student"
	"github.com/ardin2001/backend-pemilu/controllers/admin/student/repository"
	"github.com/ardin2001/backend-pemilu/models"
)

type StudentUsecaseInterface interface {
	GetAll() (*[]student.StudentResponse, error)
	GetById(id string) (*student.StudentResponse, error)
	Create(student *student.CreateStudent) error
	Update(student *student.UpdateStudent) error
	Delete(id string) error
}

type StudentUsecaseStruct struct {
	StudentRepository repository.StudentRepositoryInterface
}

func NewStudentUsecase(StudentRepository repository.StudentRepositoryInterface) StudentUsecaseInterface {
	return &StudentUsecaseStruct{
		StudentRepository: StudentRepository,
	}
}

func (su *StudentUsecaseStruct) GetAll() (*[]student.StudentResponse, error) {
	students, err := su.StudentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (su *StudentUsecaseStruct) GetById(id string) (*student.StudentResponse, error) {
	student, err := su.StudentRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (su *StudentUsecaseStruct) Create(student *student.CreateStudent) error {
	newStudent := models.Student{
		ID:  student.ID,
		NIM: student.NIM,
	}

	err := su.StudentRepository.Create(&newStudent)
	if err != nil {
		return err
	}
	return nil
}

func (su *StudentUsecaseStruct) Update(student *student.UpdateStudent) error {
	newStudent := models.Student{
		ID:  student.ID,
		NIM: student.NIM,
	}

	err := su.StudentRepository.Update(&newStudent)
	if err != nil {
		return err
	}
	return nil
}

func (su *StudentUsecaseStruct) Delete(id string) error {
	err := su.StudentRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
