package controller

import (
	"github.com/labstack/echo/v4"
	"linkshorter/logic"
	"linkshorter/models"
	"net/http"
)
var url models.Url
func CreateLink(c echo.Context)error  {
	var ep = models.Url{}
	err :=  c.Bind(&ep)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity,err.Error())
	}
	 resp, err := logic.UrlChanger(ep)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity,err.Error())
	}else {
		return c.JSON(http.StatusCreated,resp)
	}
}



