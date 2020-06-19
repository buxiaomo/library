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
	c.TplName = "search.html"
}