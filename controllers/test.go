package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Ctx.WriteString("AA11123")
}
func (c *TestController) Hello() {
	var html = "&gt;bb>aa="
	//html = template.(html)
	fmt.Println("html:", html)
	c.Ctx.WriteString(html)
}
func (c *TestController) Shou() {
	var openid = c.GetString("_openid")
	var wxid = `10004`
	var url = c.Ctx.Request.RequestURI
	fmt.Println(`zzzzzzzz` + url)
	var scope = "snsapi_userinfo"
	//c.Ctx.Redirect(302,`http://qt.sjzapps.com/wxmp/Oauth2?wxid=`+wxid+`&scope=`+scope+`&url=`+url+``)
	fmt.Println(`zzzzzzzz` + openid)
	if openid == "" {
		c.Ctx.Redirect(302, `http://qt.sjzapps.com/wxmp/Oauth2?wxid=`+wxid+`&scope=`+scope+`&url=`+url+``)
	}

	c.Ctx.WriteString(openid)
}
