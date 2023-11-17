package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=root password=root dbname=test port=10086 sslmode=disable TimeZone=Asia/Shanghai"
	Eloquent, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("postgres connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
