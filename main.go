package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"project/global"
	_ "project/routers"
)

func main() {
	defer global.Port.Close()
	beego.Run()
}
