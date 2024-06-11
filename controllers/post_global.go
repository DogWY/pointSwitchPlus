package controllers

import (
	"bytes"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
	"time"
)

type PostGlobalController struct {
	beego.Controller
}

func (c *PostGlobalController) Post() {
	m := make(map[string]any)

	addr, err := c.GetInt("addr")
	if err != nil {
		fmt.Println("获取地址失败")
	}
	global.Port.SendData(123, []byte("heart"), "udp")
	_, data, _ := global.Port.ReceiveData()
	time.Sleep(time.Second * 2)
	if bytes.Equal(data, []byte("ok")) {
		m["state"] = 1
		global.States[addr] = 1
	} else {
		m["state"] = -1
		global.States[addr] = -1
	}
	// fmt.Println("errd的值:", err)
	// global.States[addr] = -1
	// m["state"] = -1
	c.Data["json"] = m
	c.ServeJSON()
}
