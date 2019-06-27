package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var js = `<script>
		window.location='/adm/login';    
	</script>`
	c.Ctx.WriteString(js)
}
func (c *MainController) Hello() {

	c.Ctx.WriteString("ok")
}
