package db

import (
	"api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectGORM() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_graphql_sample_server"
	CONNECTION := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(DBMS, CONNECTION)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}
