package controller

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	e.POST("CreateLink", CreateLink)
	e.GET("/GetLink", GetLink)
	e.GET("/:link", Redirect)
}
