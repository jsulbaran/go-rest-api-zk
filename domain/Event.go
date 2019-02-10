package domain

type Event struct {
	ID            int    `gorm:"column:ID;primary_key"`
	UserId        string `gorm:"column:User_PIN"`
	EventDateTime string `gorm:"column:Verify_Time"`
}

func (Event) TableName() string {
	return "ATT_LOG"
}
