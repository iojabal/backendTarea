package main

import (
	"backend/handlers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	handlers.RouterInit(e)

	e.Logger.Fatal(e.Start(":80"))

}
