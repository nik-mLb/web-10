package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetQuery(c echo.Context) error {
	name := c.QueryParam("name")

	if name == ""{
		return c.String(http.StatusBadRequest, "Не введен параметр!")
	}

	test, err := srv.uc.FetchQuery(name)
	if !test && err == nil{
		return c.String(http.StatusBadRequest, "Запись не добавлена в БД!")
	} else if (!test && err != nil){
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Hello, "+name+"!")
}

func (srv *Server) PostQuery(c echo.Context) error {
	name := c.QueryParam("name")
	if name == ""{
		return c.String(http.StatusBadRequest, "Не введен параметр!")
	}

	test, err := srv.uc.FetchQuery(name)
	if test && err == nil{
		return c.String(http.StatusBadRequest, "Запись уже добавлена в БД!")
	}

	err = srv.uc.SetQuery(name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "Добавили запись!")
}