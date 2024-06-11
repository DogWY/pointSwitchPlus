package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
)

type DisplayController struct {
	beego.Controller
}

// Post 距离展示控制层, 接收post请求, 参数为lora模块地址, 返回最近 length 获取到的距离以及对应时间
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
