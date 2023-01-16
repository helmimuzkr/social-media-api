package main

import (
	"log"
	"social-media-app/config"

	"github.com/labstack/echo"
)

func main() {
	c := config.GetConfig()
	db := config.OpenDB(c)

	e := echo.New()
	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
