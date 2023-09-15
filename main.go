package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
	_ "pointSwitch/routers"
)

func main() {
	defer global.Port.Close()
	beego.Run()
}
