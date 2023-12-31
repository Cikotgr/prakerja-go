package repository

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/admin"
	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	LoginAdmin(username, password string) (*admin.AdminResponse, error)
	GetById(id string) (*admin.AdminDetailResponse, error)
}

type AdminRepositoryStruct struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) AdminRepositoryInterface {
	return &AdminRepositoryStruct{
		DB: DB,
	}
}

func (ar *AdminRepositoryStruct) GetById(id string) (*admin.AdminDetailResponse, error) {
	var admin admin.AdminDetailResponse
	err := ar.DB.Table("admins").Where("id = ? ", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (ar *AdminRepositoryStruct) LoginAdmin(username, password string) (*admin.AdminResponse, error) {
	var admin admin.AdminResponse
	err := ar.DB.Table("admins").Where("username = ? AND password = ?", username, password).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
