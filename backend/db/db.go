package db

import (
	"fmt"

	"gorm.io/driver/mysql" 
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	var err error

	dsn := "root:root@tcp(localhost:3306)/forum?parseTime=true"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Database connected and tables migrated successfully")
}
