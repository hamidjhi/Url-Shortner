package main

import (
	"linkshorter/controller"
	"linkshorter/db"

"github.com/labstack/echo/v4"
)
func main()  {

	 e := echo.New()
	 db.ConnectToRedis()
	 controller.Routes(e)
	e.Start("127.0.0.1:1234")

}
