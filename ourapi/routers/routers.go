package routers

import (
	"github.com/labstack/echo"
	"github.com/piggyman007/golang-training-intermediate/ourapi/users"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", users.GetUsers)

	return e
}
