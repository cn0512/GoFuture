package main

/*
	echo web frame curd operator
*/

import (
	"log"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Cfg *SvcCfg

func init() {
	Cfg = &SvcCfg{}
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(4)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/remotelog/list", GetLogList)
	e.GET("/remotelog/list/:id", GetLogData)
	e.POST("/remotelog/log/list", CreateLogList)
	e.POST("/remotelog/log/detail", CreateLogDetail)
	/*
		//save routers to file
		data, err := json.MarshalIndent(e.Routes(), "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		ioutil.WriteFile("routes.json", data, 0644)
	*/
	//load cfg
	Parse("./echo.yaml", Cfg)
	log.Println(Cfg)
	if strings.EqualFold(Cfg.Http, "") {
		log.Fatal("cfg load err")
	}

	//init pub
	InitPub(Cfg.Nsq_tcp)

	// Start server
	e.Logger.Fatal(e.Start(Cfg.Http)) //":8080"
}
