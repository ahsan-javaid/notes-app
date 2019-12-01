package models

import (
"github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)
var DB *gorm.DB
func Connect() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=root sslmode=disable")
	if err != nil {
		panic(err)
	} else  {
		log.Printf("Connected to database")
		db.LogMode(true)
		db.AutoMigrate(&User{})
	}
    DB = db
}