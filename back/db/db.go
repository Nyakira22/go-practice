package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	// if os.Getenv("GO_ENV") == "dev" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	DBMS := "mysql"
	USER := "root"
	PASS := "root_password"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "test"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connected!")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
