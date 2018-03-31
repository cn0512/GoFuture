package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

//----------
// Handlers
//----------

type Listbuf struct {
	Id string `json:"id"`
}

type LogDetail struct {
	Id  string `json:"id"`
	Buf string `json:"buf"`
}

func GetList() []byte {
	listbuf := &Listbuf{"helloworld"}
	buf, _ := json.Marshal(listbuf)
	return buf
}

func GetLogList(c echo.Context) error {

	listbuf := GetList()
	return c.JSON(http.StatusOK, listbuf)
}

func GetLogData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, id)
}
