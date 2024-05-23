package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
)

type PostGlobalController struct {
	beego.Controller
}

func (c *PostGlobalController) Post() {
	c.Ctx.Output.ContentType("application/json")

	m := make(map[string]any)

	// 从请求中获取参数的字符串值
	//addrStr := c.GetString("addr")
	addrStr := c.Ctx.Request.PostFormValue("addr")
	fmt.Println("addr 参数值（字符串）:", addrStr)

	addr, err := c.GetInt("addr")
	if err != nil {
		fmt.Println("获取地址失败")
	}
	//global.Port.SendData(123, []byte("heart"), "udp")
	//_, data, _ := global.Port.ReceiveData()
	//time.Sleep(time.Second * 2)
	//if bytes.Equal(data, []byte("ok")) {
	//	m["state"] = 1
	//	global.States[addr] = 1
	//} else {
	//	m["state"] = -1
	//	global.States[addr] = -1
	//}

	global.States[addr] = -1
	m["state"] = -1
	c.Data["json"] = m
	c.ServeJSON()
}
