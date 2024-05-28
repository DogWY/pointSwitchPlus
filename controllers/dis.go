package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
	// "strconv"
	// "time"
)

type DisController struct {
	beego.Controller
}

// func (c *DisController) Get() {
// 	c.Data["addrs"] = global.Addrs
// 	c.TplName = "dis.html"
// }

func (c *DisController) Post() {
	c.Ctx.Output.ContentType("application/json")
	m := make(map[string]any)
	addr, err := c.GetInt("addr")
	if err != nil {
		fmt.Println("获取地址失败")
	}
	// global.Port.SendData(123, []byte("dis"), "udp")
	// _, data, err := global.Port.ReceiveData()
	// if err == nil {
	// 	tmp, _ := strconv.Atoi(string(data))
	// 	for i := 0; i < global.Length-1; i++ {
	// 		global.MDis[addr][i] = global.MDis[addr][i+1]
	// 		global.MTimes[addr][i] = global.MTimes[addr][i+1]
	// 	}
	// 	global.MDis[addr][global.Length-1] = tmp
	// }
	// global.MTimes[addr][global.Length-1] = time.Now().Format("15:04:05")
	// global.MDis[123] = [1,2,3,4,5]
	// global.MTimes[123] = {"0s","1s","2s","3s","4s"}
	global.MTimes[90][0] = "0s"
	global.MTimes[90][1] = "1s"
	global.MTimes[90][2] = "2s"
	global.MTimes[90][3] = "3s"
	global.MTimes[90][4] = "4s"
	global.MTimes[121][0] = "0s"
	global.MTimes[121][1] = "1s"
	global.MTimes[121][2] = "2s"
	global.MTimes[121][3] = "3s"
	global.MTimes[121][4] = "4s"
	global.MTimes[123][0] = "0s"
	global.MTimes[123][1] = "1s"
	global.MTimes[123][2] = "2s"
	global.MTimes[123][3] = "3s"
	global.MTimes[123][4] = "4s"
	global.MDis[90][0] = 0
	global.MDis[90][1] = 1
	global.MDis[90][2] = 0
	global.MDis[90][3] = 0
	global.MDis[90][4] = 1
	global.MDis[121][0] = 1
	global.MDis[121][1] = 0
	global.MDis[121][2] = 0
	global.MDis[121][3] = 0
	global.MDis[121][4] = 1
	global.MDis[123][0] = 0
	global.MDis[123][1] = 1
	global.MDis[123][2] = 0
	global.MDis[123][3] = 1
	global.MDis[123][4] = 1

	m["dis"] = global.MDis[addr]
	m["times"] = global.MTimes[addr]
	c.Data["json"] = m
	c.ServeJSON()
}
