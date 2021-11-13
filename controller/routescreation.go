package controller

import (
	"github.com/labstack/echo/v4"
	"linkshorter/logic"
	"linkshorter/models"
	"net/http"
)

var url models.Url

func CreateLink(c echo.Context) error {
	var ep = models.Url{}
	err := c.Bind(&ep)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	resp, err := logic.UrlChanger(ep)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		return c.JSON(http.StatusCreated, resp)
	}
}
func GetLink(c echo.Context) error {

	shortLink := c.QueryParam("short_link")
	res, err := logic.UrlReturn(shortLink)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

func Redirect(c echo.Context) error {
	shortLink := c.Param("link")
	res, err := logic.UrlReturn(shortLink)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		//return c.HTML(http.StatusOK, fmt.Sprintf("<a href='%s'> click here </a>", res.LongUrl))
		return c.Redirect(http.StatusSeeOther, res.LongUrl)
	}
}
