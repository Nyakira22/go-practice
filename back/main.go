package main

import (
	"github.com/akiradomi/workspace/go-practice/back/controller"
	"github.com/akiradomi/workspace/go-practice/back/db"
	"github.com/akiradomi/workspace/go-practice/back/repository"
	"github.com/akiradomi/workspace/go-practice/back/router"
	"github.com/akiradomi/workspace/go-practice/back/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8000"))
}

// func sqlConnect() (database *gorm.DB) {
// 	DBMS := "mysql"
// 	USER := "root"
// 	PASS := "root_password"
// 	PROTOCOL := "tcp(0.0.0.0:3306)"
// 	DBNAME := "test"

// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

// 	count := 0
// 	db, err := gorm.Open(DBMS, CONNECT)
// 	if err != nil {
// 		for {
// 			if err == nil {
// 				fmt.Println("")
// 				break
// 			}
// 			fmt.Print(".")
// 			time.Sleep(time.Second)
// 			count++
// 			if count > 180 {
// 				fmt.Println("")
// 				fmt.Println("DB接続失敗")
// 				panic(err)
// 			}
// 			db, err = gorm.Open(DBMS, CONNECT)
// 		}
// 	}
// 	fmt.Println("DB接続成功")

// 	return db
// }
