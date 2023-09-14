package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"project/global"
	"project/myjpeg"
)

type ImgController struct {
	beego.Controller
}

func (c *ImgController) Get() {
	c.Data["msg"] = ""
	c.Data["addrs"] = global.Addrs
	c.TplName = "system_img.html"
}

func (c *ImgController) Post() {
	addr, err := c.GetInt("addr")
	if err != nil {
		fmt.Println("获取地址失败")
	}
	way, err := c.GetInt("transway")
	if err != nil {
		fmt.Println("获取方式失败")
	}
	if way == 1 {
		global.Port.SendData(addr, []byte("img1"), "udp")
	} else {
		global.Port.SendData(addr, []byte("img0"), "udp")
		_, data, _ := global.Port.ReceiveData()
		myjpeg.Decode("./static/img/123_monitor.jpeg", data)
	}
	m := make(map[string]any)
	m["message"] = "初始化已经完成"
	c.Data["json"] = m
	c.ServeJSON()
}
