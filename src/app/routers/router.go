package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
	
)

func init() {
	beego.Info("Loading router...")
	
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.LoginController{}, "get:ShowLoginPage")
	beego.Router("/login/:provider", &controllers.LoginController{}, "get:Authenticate")
	beego.Router("/timlogin", &controllers.LoginController{}, "post:TimAuthenticate")
	beego.Router("/auth/:provider/callback", &controllers.LoginController{}, "get:Validate")
	beego.Router("/secure", &controllers.SecureContent{})
}
