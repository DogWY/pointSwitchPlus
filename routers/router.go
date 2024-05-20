package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"pointSwitch/controllers"
)

func init() {
	beego.SetStaticPath("/favicon.ico", "/static/favicon.ico")
	beego.Router("/img", &controllers.ImgController{})
	beego.Router("/upload_img", &controllers.UploadController{})
	beego.Router("/init_img", &controllers.ImgController{})

	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/system", &controllers.SystemController{})
	beego.Router("/init_system", &controllers.SystemController{})

	beego.Router("/system_conf", &controllers.ConfController{})

	beego.Router("/dis", &controllers.DisController{})
	beego.Router("/display", &controllers.DisplayController{})

	beego.Router("/get_global", &controllers.GetGlobalController{})
}
