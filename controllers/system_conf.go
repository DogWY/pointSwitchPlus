package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
	"pointSwitch/services"
	"strings"
)

type ConfController struct {
	beego.Controller
}

func (c *ConfController) Get() {
	c.Data["addrs"] = global.Addrs
	c.TplName = "system_conf.html"
}

func (c *ConfController) Post() {
	newAddrs := c.Ctx.Request.PostFormValue("addrs")
	addrs := strings.Split(newAddrs, ",")
	msg := c.Ctx.Request.PostFormValue("msg")
	if msg == "增加" {
		services.AddAddrs(addrs)
	} else {
		services.RemoveAddrs(addrs)
	}
	c.Data["addrs"] = global.Addrs
	c.TplName = "system_conf.html"
}
