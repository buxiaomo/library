package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	keyword := c.Input().Get("keyword")
	fmt.Println(keyword)
	c.Data["Title"] = beego.AppConfig.DefaultString("title","书库")
	c.TplName = "search.html"
}