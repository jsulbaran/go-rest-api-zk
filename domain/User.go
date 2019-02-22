package domain

import (
	"errors"
	"net/http"
)

type User struct {
	InternalId  int        `json:"id" gorm:"column:ID;primary_key"`
	UserPin     string     `json:"userId" gorm:"column:User_PIN"`
	Name        string     `json:"name" gorm:"column:Name"`
	Rut         string     `json:"cod,omitempty" gorm:"-"`
	Card        string     `json:"card_number,omitempty" gorm:"column:Main_Card"`
	CompanyId   int        `json:"company_id" gorm:"column:Company_ID"`
	RoleNumber  int        `json:"-" gorm:"column:Privilege"`
	RoleString  string     `json:"role,omitempty" gorm:"-"`
	ErrorString string     `json:"error_desc,omitempty" gorm:"-"`
	Biometric   []Template `gorm:"foreignkey:UserId;association_foreignkey:InternalId" json:"biometric,omitempty"`
	ModifyTime  string     `gorm:"column:MODIFY_TIME" json:"-"`
	//Events  []Event `gorm:"foreignkey:UserId;association_foreignkey:UserPin" json:"events"`
}

//type UserPayload struct {
//	*User
//	Role string `json:"role"`
//}
type UserRequest struct {
	*User
	//Info *UserPayload `json:"user,omitempty"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.User == nil {
		return errors.New("missing User")
	}

	//// a.User is nil if no Userpayload fields are sent in the request. In this app
	//// this won't cause a panic, but checks in this Bind method may be required if
	//// a.User or futher nested fields like a.User.Name are accessed elsewhere.
	//
	//// just a post-process after a decode..
	//a.ProtectedID = ""                                 // unset the protected ID
	//a.Article.Title = strings.ToLower(a.Article.Title) // as an example, we down-case
	return nil
}

func (User) TableName() string {
	return "USER_INFO"
}
