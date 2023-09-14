package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"project/global"
)

type DisplayController struct {
	beego.Controller
}

func (c *DisplayController) Post() {
	addr, err := c.GetInt("addr")
	if err != nil {
		fmt.Println("获取地址失败")
	}
	m := make(map[string]any)
	m["dis"] = global.MDis[addr]
	m["times"] = global.MTimes[addr]
	c.Data["json"] = m
	c.ServeJSON()
}
