package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"strconv"
)

func (srv *Server) GetCounter (c echo.Context) error {
	msg, err := srv.uc.FetchCount()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Счетчик: " + strconv.Itoa(msg))
}

func (srv *Server) PostCounter (c echo.Context) error{
	input := struct {
		Msg int `json:"msg"`
	}{}

	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	err = srv.uc.SetCounter(input.Msg)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Изменили счетчик!")
}