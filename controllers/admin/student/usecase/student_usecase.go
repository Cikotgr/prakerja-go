package usecase

import (
	"net/http"

	"github.com/ardin2001/backend-pemilu/controllers/admin/student"
	"github.com/ardin2001/backend-pemilu/controllers/admin/student/repository"
	"github.com/ardin2001/backend-pemilu/models"
	"github.com/pkg/errors"
)

type StudentUsecaseInterface interface {
	GetAll(nim string) (*[]student.StudentResponse, int, error)
	GetById(id string) (*student.StudentResponse, int, error)
	Create(student *student.CreateStudent) (int, error)
	Update(student *student.UpdateStudent) (int, error)
	Delete(id string) (int, error)
}

type StudentUsecaseStruct struct {
	StudentRepository repository.StudentRepositoryInterface
}

func NewStudentUsecase(StudentRepository repository.StudentRepositoryInterface) StudentUsecaseInterface {
	return &StudentUsecaseStruct{
		StudentRepository: StudentRepository,
	}
}

func (su *StudentUsecaseStruct) GetAll(nim string) (*[]student.StudentResponse, int, error) {
	students, err := su.StudentRepository.GetAll(nim)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to get all students data")
	}
	return students, http.StatusOK, nil
}

func (su *StudentUsecaseStruct) GetById(id string) (*student.StudentResponse, int, error) {
	student, err := su.StudentRepository.GetById(id)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to get student detail data")
	}
	return student, http.StatusOK, nil
}

func (su *StudentUsecaseStruct) Create(student *student.CreateStudent) (int, error) {
	newStudent := models.Student{
		ID:     student.ID,
		NIM:    student.NIM,
		RoleId: student.RoleId,
	}

	err := su.StudentRepository.Create(&newStudent)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to create student data")
	}
	return http.StatusOK, nil
}

func (su *StudentUsecaseStruct) Update(student *student.UpdateStudent) (int, error) {
	newStudent := models.Student{
		ID:  student.ID,
		NIM: student.NIM,
	}

	err := su.StudentRepository.Update(&newStudent)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to update student data")
	}
	return http.StatusOK, nil
}

func (su *StudentUsecaseStruct) Delete(id string) (int, error) {
	err := su.StudentRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to delete student data")
	}
	return http.StatusOK, nil
}
