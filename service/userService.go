package service

import (
	"RestService/domain"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(orm *gorm.DB) *UserService {
	return &UserService{db: orm}
}

func (service UserService) GetUsersWithTemplates() []domain.User {
	var users []domain.User
	service.db.Debug().Preload("Biometric").Find(&users)
	//service.db.Preload("Biometric").Find(&users)
	return users
}
func (service UserService) GetUsers() domain.User {
	var users domain.User
	service.db.Find(&users)
	return users
}
