package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct{
	beego.Controller
}

func (this *MainController) Get(){
	this.Data["Website"] = "shaniu.site"
	this.Data["Email"] = "shenhua3@huawei.com"
	this.Ctx.WriteString("Welcome shaniu.site")
}