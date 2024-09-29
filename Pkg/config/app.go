package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("postgres", "postgresql://management_erp:nq6-jRS0qHZIA6GuKuZWxQ@zinc-ogress-7322.6xw.aws-ap-southeast-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
