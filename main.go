package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Res struct {
	Status   string
	NodeIp   string
	NodeName string
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		nodeIp := os.Getenv("NODE_IP")
		nodeName := os.Getenv("NODE_NAME")
		return c.JSON(http.StatusOK, Res{Status: "OK", NodeIp: nodeIp, NodeName: nodeName})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

}
