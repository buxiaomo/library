package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = beego.AppConfig.DefaultString("title","书库")
	c.TplName = "index.html"
}