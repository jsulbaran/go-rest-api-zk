package domain

type User struct {
	InternalId int        `json:"id" gorm:"column:ID;primary_key"`
	UserPin    string     `json:"userId" gorm:"column:User_PIN"`
	Name       string     `json:"name" gorm:"column:Name"`
	Rut        string     `json:"cod,omitempty"`
	Card       string     `json:"card_number,omitempty" gorm:"column:Main_Card"`
	CompanyId  int        `json:"company_id" gorm:"column:Company_ID"`
	RoleNumber int        `json:"-" gorm:"column:Privilege"`
	RoleString string     `json:"role,omitempty"`
	Biometric  []Template `gorm:"foreignkey:UserId;association_foreignkey:InternalId" json:"biometric,omitempty"`
	//Events  []Event `gorm:"foreignkey:UserId;association_foreignkey:UserPin" json:"events"`
}

func (User) TableName() string {
	return "USER_INFO"
}
