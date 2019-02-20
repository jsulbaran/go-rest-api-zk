package service

import (
	"RestService/domain"
	"RestService/util"
	"github.com/dustin/go-humanize"
	"github.com/jinzhu/gorm"
	"strings"
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
func (service UserService) GetUsers() []domain.User {
	var users []domain.User
	//service.db.Find(&users)
	//service.db.Debug().Preload("Biometric").Find(&users)
	service.db.Preload("Biometric").Find(&users)
	return users
}

func (service UserService) GetUserById(id int) domain.User {
	var user domain.User
	//service.db.Where(&domain.User{UserPin: id}).First(&user)
	service.db.Preload("Biometric").First(&user, id)
	if user.InternalId > 0 {
		user = completeUserData(user)
	}
	return user
}

func (service UserService) GetUserByPin(pin string) domain.User {
	var user domain.User
	service.db.Preload("Biometric").Where(&domain.User{UserPin: pin}).First(&user)
	if user.InternalId > 0 {
		user = completeUserData(user)
	}
	return user
}
func completeUserData(user domain.User) domain.User {
	if user.RoleNumber == 0 {
		user.RoleString = "user"
	} else {
		user.RoleString = "admin"
	}
	rutInt := util.StringToInt(user.UserPin)
	digit := util.CalcularDigitoRut(rutInt)
	rutStr := humanize.Comma(int64(rutInt))
	rutStr = strings.Replace(rutStr, ",", ".", -1)
	user.Rut = rutStr + "-" + string(digit)
	return user
}
