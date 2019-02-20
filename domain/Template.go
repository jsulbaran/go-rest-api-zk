package domain

type Template struct {
	ID     int    `gorm:"column:ID;primary_key" json:"-"`
	UserId int    `gorm:"column:pin" json:"-"`
	Size   int    `gorm:"column:size" json:"templateSize"`
	Data   []byte `gorm:"column:template" json:"data,omitempty"`
}

func (Template) TableName() string {
	return "fptemplate10"
}
