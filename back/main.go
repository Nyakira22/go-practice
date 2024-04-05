package main

import (
	"log"

	"github.com/akiradomi/workspace/go-practice/back/config"
	"github.com/akiradomi/workspace/go-practice/back/controller"
	"github.com/akiradomi/workspace/go-practice/back/db"
	"github.com/akiradomi/workspace/go-practice/back/repository"
	"github.com/akiradomi/workspace/go-practice/back/router"
	"github.com/akiradomi/workspace/go-practice/back/usecase"
	"github.com/akiradomi/workspace/go-practice/back/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.LoggingSettings(config.Config.Logging)
	log.Println("log test")
	//DBインスタンス
	db := db.NewDB()
	//user_repositoryのインスタンス化
	userRepository := repository.NewUserRepository(db)
	//user_usecaseのインスタンス化
	userUsecase := usecase.NewUserUsecase(userRepository)
	//user_controllerのインスタンス化
	userController := controller.NewUserController(userUsecase)
	//routerのインスタンス化
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8000"))
}
