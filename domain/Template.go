package domain

type Template struct {
	ID     int `gorm:"column:ID;primary_key"`
	UserId int `gorm:"column:pin" json:"-"`
	Size   int `gorm:"column:size" json:"templateSize"`
}

func (Template) TableName() string {
	return "fptemplate10"
}
