package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique_index" db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Hobby    string `db:"hobby" json:"hobby"`
}
