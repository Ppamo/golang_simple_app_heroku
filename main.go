package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

var (
	SERVER_PORT=os.Getenv("PORT")
	FLAVOUR=os.Getenv("FLAVOUR")
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error{
		fmt.Printf("> ping\n")
		return c.String(http.StatusOK, "{\"status\":\"pong\", \"flavour\": \"" + FLAVOUR + "\"}")
	})
	e.Logger.Fatal(e.Start(":" + SERVER_PORT))
}
