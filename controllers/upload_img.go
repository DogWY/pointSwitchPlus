package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Post() {
	tofile := "./static/img/7.jpeg"
	err := c.SaveToFile("image", tofile)
	if err != nil {
		fmt.Println("4G接受文件错误")
	}
}
