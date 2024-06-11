package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
	"strconv"
	"time"
)

type DisController struct {
	beego.Controller
}

func (c *DisController) Get() {
	c.Data["addrs"] = global.Addrs
	c.TplName = "dis.html"
}

func (c *DisController) Post() {
	addr, err := c.GetInt("addr")
	if err != nil {
		fmt.Println("获取地址失败")
	}
	global.Port.SendData(123, []byte("dis"), "udp")
	_, data, err := global.Port.ReceiveData()
	fmt.Println(string(data))
	if err == nil {
		tmp, _ := strconv.ParseFloat(string(data), 64)
		fmt.Println(tmp)
		for i := 0; i < global.Length-1; i++ {
			global.MDis[addr][i] = global.MDis[addr][i+1]
			global.MTimes[addr][i] = global.MTimes[addr][i+1]
		}
		global.MDis[addr][global.Length-1] = tmp
	}
	global.MTimes[addr][global.Length-1] = time.Now().Format("15:04:05")
	m := make(map[string]any)
	fmt.Println(global.MDis[addr])
	m["dis"] = global.MDis[addr]
	m["times"] = global.MTimes[addr]
	c.Data["json"] = m
	c.ServeJSON()
}
