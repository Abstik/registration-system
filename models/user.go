package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"type:varchar(64);not null;" json:"name"`
	Password   string `gorm:"type:varchar(64);not null" json:"-"`
	Major      string `gorm:"type:varchar(64);not null" json:"major"` //专业
	Phone      string `gorm:"type:varchar(64);not null" json:"phone"`
	QQ         string `gorm:"type:varchar(64);not null" json:"qq"`
	Email      string `gorm:"type:varchar(64);not null;unique" json:"email"`
	Direction  string `gorm:"type:varchar(64);not null" json:"direction"`
}
