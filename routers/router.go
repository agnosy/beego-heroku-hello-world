package routers

import (
	"github.com/adamar/beego-heroku-hello-world/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
