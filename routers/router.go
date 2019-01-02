package routers

import (
	"github.com/astaxie/beego"
	"webandsocket/controllers"
)

func init() {
    beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleRegister")
}
