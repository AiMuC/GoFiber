package model

type Users struct {
	Id   int32  `gorm:"column:id;type:INT(11);AUTO_INCREMENT;NOT NULL"`
	Name string `gorm:"column:name;type:VARCHAR(10);NOT NULL"`
}
