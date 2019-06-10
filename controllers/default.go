package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {

	this.Layout = "layouts/layout.tpl"
	this.Data["assets_domain"] = beego.AppConfig.String("assets_domain")
	this.Data["site_title"] = beego.AppConfig.String("site_title")

}

func (this *MainController) Get() {

	this.TplName = "index.tpl"

}
