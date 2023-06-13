package datasource

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConfigData() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //&gom.Config: cung cấp cấu hình

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)
	return db
}
