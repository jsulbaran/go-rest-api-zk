package domain

type Template struct {
	ID         int    `gorm:"column:ID;primary_key" json:"-"`
	FingerID   int    `gorm:"column:fingerid;default:1" json:"-" default:"1"`
	Valid      int    `gorm:"column:valid;default:1" json:"-"`
	UserId     int    `gorm:"column:pin" json:"-"`
	Size       int    `gorm:"column:size" json:"-"`
	Data       []byte `gorm:"column:template" json:"data,omitempty"`
	ModifyTime string `gorm:"column:MODIFY_TIME" json:"-"`
}

func (Template) TableName() string {
	return "fptemplate10"
}
