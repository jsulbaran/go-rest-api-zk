package domain

type Event struct {
	ID            int    `gorm:"column:ID;primary_key"`
	UserId        string `gorm:"column:User_PIN"`
	EventDateTime string `gorm:"column:Verify_Time"`
	VerifyType    string `gorm:"column:Verify_Type"`
	Status        int    `gorm:"column:Status"`
}

func (Event) TableName() string {
	return "ATT_LOG"
}
