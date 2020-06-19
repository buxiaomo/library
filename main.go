package main

import (
	"github.com/astaxie/beego"
	_ "github.com/buxiaomo/library/models"
	_ "github.com/buxiaomo/library/routers"
)

func main() {
	beego.Run()
}