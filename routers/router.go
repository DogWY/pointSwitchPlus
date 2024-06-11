package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"pointSwitch/controllers"
)

func init() {

	// 创建 CORS 中间件
	corsMw := cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},                                       // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 允许暴露的响应头
		AllowCredentials: true,                                                // 是否允许发送凭证信息（如 Cookie）
	})

	// 插入 CORS 中间件
	beego.InsertFilter("*", beego.BeforeRouter, corsMw)

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
	beego.Router("/post_global", &controllers.PostGlobalController{})

}
