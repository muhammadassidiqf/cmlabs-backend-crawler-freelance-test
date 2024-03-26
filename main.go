package main

import (
	"github.com/labstack/echo/v4"
	"crawler-website/handler"
)

func main()  {
	app := echo.New()

	crawlerhandler := handler.InitCrawlerHandler()

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Selamat datang di api Crawler")
	})

	app.GET("/cmlabs", crawlerhandler.FetchData("https://cmlabs.co"))
	app.GET("/sequence", crawlerhandler.FetchData("https://www.sequence.day"))
	app.GET("/posibel", crawlerhandler.FetchData("https://posibel.co"))
	app.GET("/bwa", crawlerhandler.FetchData("https://buildwithangga.com/"))

	app.Start(":7777")
}
