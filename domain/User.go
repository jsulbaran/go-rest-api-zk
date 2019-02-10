package domain

type User struct {
	InternalId int        `json:"internalId" gorm:"column:ID;primary_key"`
	UserPin    string     `json:"userId" gorm:"column:User_PIN"`
	Name       string     `json:"name" gorm:"column:Name"`
	Biometric  []Template `gorm:"foreignkey:UserId;association_foreignkey:InternalId" json:"biometric"`
}

func (User) TableName() string {
	return "USER_INFO"
}
