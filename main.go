package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

var (
	SERVER_PORT=os.Getenv("SERVER_PORT")
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error{
		fmt.Printf("> ping\n")
		return c.String(http.StatusOK, "{\"status\":\"pong\"}")
	})
	e.Logger.Fatal(e.Start(":" + SERVER_PORT))
}
