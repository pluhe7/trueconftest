package main

import (
	"github.com/pluhe7/trueconftest/handler"
	"github.com/pluhe7/trueconftest/model"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	h := handler.NewHandler(map[int]*model.User{}, 1)
	h.Register(e)

	e.Start(":8080")
}
