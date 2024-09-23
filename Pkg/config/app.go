package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("postgres", "host=localhost port=5432 dbname=issue user=postgres password=<your_db_password> connect_timeout=10 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
