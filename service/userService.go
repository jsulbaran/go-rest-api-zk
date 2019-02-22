package service

import (
	"RestService/domain"
	"RestService/util"
	"github.com/dustin/go-humanize"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type UserService struct {
	db *gorm.DB
}

const ADMIN_ROLE_CODE = 120
const USER_ROLE_CODE = 0
const USER_ROLE_STRING = "user"
const ADMIN_ROLE_STRING = "admin"

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

func (service UserService) DeleteUserAndTemplateById(id int) bool {
	user := domain.User{InternalId: id}
	//delete templates
	service.deleteTemplateByUserId(id)
	service.db.Delete(&user)
	return true
}

func getRoleStringFromCode(code int) string {
	if code == ADMIN_ROLE_CODE {
		return ADMIN_ROLE_STRING
	}
	return USER_ROLE_STRING
}

func getRoleCodeFromString(roleString string) int {
	if roleString == ADMIN_ROLE_STRING {
		return ADMIN_ROLE_CODE
	}
	return USER_ROLE_CODE
}

func completeUserData(user domain.User) domain.User {
	user.RoleString = getRoleStringFromCode(user.RoleNumber)
	rutInt := util.StringToInt(user.UserPin)
	digit := util.CalcularDigitoRut(rutInt)
	rutStr := humanize.Comma(int64(rutInt))
	rutStr = strings.Replace(rutStr, ",", ".", -1)
	user.Rut = rutStr + "-" + string(digit)
	return user
}
func (service UserService) deleteTemplateByUserId(id int) {
	service.db.Where("pin = ?", id).Delete(domain.Template{})

}

func (service UserService) IsValidUser(user domain.User) bool {

	return true
}

func (service UserService) UpdateUserById(userId int, user domain.User) domain.User {
	var dbUser domain.User
	service.db.First(&dbUser, userId)
	dbUser.Rut = user.Rut
	dbUser.UserPin = user.UserPin
	dbUser.Name = user.Name
	dbUser.CompanyId = user.CompanyId
	dbUser.RoleString = user.RoleString

	modifiedTime := time.Now().Format("20060102150405")
	dbUser.ModifyTime = modifiedTime

	//service.db.Where("pin = ?", userId).First(domain.Template{})
	//
	service.db.Where("pin = ?", userId).Delete(domain.Template{})

	if len(user.Biometric) >= 1 {
		dbUser.Biometric = []domain.Template{{Valid: 1, FingerID: 1, Size: len(user.Biometric[0].Data), Data: user.Biometric[0].Data, ModifyTime: modifiedTime}}

		//	templateId := 0
		//	var template domain.Template
		//	service.db.Where("pin = ?", userId).First(&template)
		//	if template.ID > 0 {
		//		templateId = template.ID
		//	}
		//	dbUser.Biometric = []domain.Template{{Valid: 1, FingerID: 1, Size: len(user.Biometric[0].Data), Data: user.Biometric[0].Data, ModifyTime: modifiedTime, ID: templateId}}
		//} else {
		//	service.db.Where("pin = ?", userId).Delete(domain.Template{})
	}
	dbUser.Card = user.Card
	service.db.Debug().Preload("Biometric").Save(dbUser)
	return dbUser
}

func (service UserService) CreateNewUser(user domain.User) domain.User {
	var dbUser domain.User
	dbUser.Rut = user.Rut
	dbUser.UserPin = user.UserPin
	dbUser.Name = user.Name
	dbUser.CompanyId = user.CompanyId
	dbUser.RoleNumber = getRoleCodeFromString(user.RoleString)
	modifiedTime := time.Now().Format("20060102150405")
	dbUser.ModifyTime = modifiedTime
	if len(user.Biometric) >= 1 {
		dbUser.Biometric = []domain.Template{{Valid: 1, FingerID: 1, Size: len(user.Biometric[0].Data), Data: user.Biometric[0].Data, ModifyTime: modifiedTime}}
	}
	dbUser.Card = user.Card
	service.db.NewRecord(dbUser) // => returns `true` as primary key is blank

	if err := service.db.Debug().Preload("Biometric").Create(&dbUser).Error; err != nil {
		return domain.User{InternalId: -1, ErrorString: err.Error()}
	}
	//service.db.Debug().Preload("Biometric").Save(dbUser)

	return completeUserData(dbUser)
}
