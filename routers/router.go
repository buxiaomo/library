package routers

import (
	"github.com/buxiaomo/library/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
    beego.Router("/search", &controllers.SearchController{})
	beego.Router("/upload", &controllers.UploadController{})
}