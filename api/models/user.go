package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique_index" json:"name"`
	Password string `json:"password"`
	Hobby    string `json:"hobby"`
}
