package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
)

type GetGlobalController struct {
	beego.Controller
}

func (c *GetGlobalController) Get() {
	c.Ctx.Output.ContentType("application/json")

	response := map[string]interface{}{
		"Addrs":  global.Addrs,
		"MDis":   global.MDis,
		"MTimes": global.MTimes,
		"States": global.States,
	}

	c.Data["json"] = response
	c.ServeJSON()
}
