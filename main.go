package main

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) HUI() {
	c.Ctx.WriteString("hello world")
}

func (c *MainController) HUI2() {
	c.Ctx.WriteString("Я твоего шо?")
}

func main() {
	web.Router("/", &MainController{}, "get:HUI")
	web.Router("/hui", &MainController{}, "get:HUI2")
	web.Run()
}
