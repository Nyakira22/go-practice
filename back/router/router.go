package router

import (
	"github.com/akiradomi/workspace/go-practice/back/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.UserControllerInterface) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	return e
}
