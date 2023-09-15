package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/global"
	"pointSwitch/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.Ctx.Request.PostFormValue("username")
	password := c.Ctx.Request.PostFormValue("password")
	var user models.User
	err := global.DB.Model(models.User{}).Where("Name = ?", username).First(&user).Error
	if err != nil {
		global.Logger.Info("login fail： wrong username",
			"username", username,
		)
		c.Data["msg"] = "用户名不存在"
		c.TplName = "login.html"
		return
	}
	if user.Pwd != password {
		global.Logger.Infow("login fail： wrong password",
			"username", username,
		)
		c.Data["msg"] = "密码不正确"
		c.TplName = "login.html"
		return
	}
	global.Logger.Info("login success",
		"username", username,
	)
	c.Data["msg"] = "登陆成功"
	c.Data["addrs"] = global.Addrs
	c.Data["states"] = global.States
	c.TplName = "system.html"
}
