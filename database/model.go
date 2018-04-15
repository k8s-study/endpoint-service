package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	// _ “github.com/go-sql-driver/mysql”
)

type Endpoint struct {
	ID uint `json:"id"`
	UserId int `json:"userId"`
	Url string `json:"url"`
	Port int `json:"port"`
	Interval int `json:"interval"`
}

func Init() (db *gorm.DB, err error){

	db, err = gorm.Open("sqlite3", "./gorm.db")
	return
}


