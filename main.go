package main

import (
	"fmt"
	"origin/conf"
	"origin/controllers"
	_ "origin/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/kardianos/service"
	adm "github.com/sjzdlm/adm/conf"
	"github.com/sjzdlm/db"

	//dbtool "github.com/sjzdlm/dbtool/conf"
	balong "github.com/sjzdlm/balong/conf"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	//启动服务
	go p.run()
	if beego.AppConfig.String("service") != "1" {
		//初始化窗体
		//mw := NewWindow()
		//mw.SetX(800)
		//mw.SetY(200)
		//mw.init()
		//mw.AddNotifyIcon()
		//mw.Run()
	}

	return nil
}

func (p *program) run() {
	// 重置模板路径
	//beego.SetViewsPath(GetAPPRootPath() + "/views/")

	db.InitX()
	adm.InitRouter()
	adm.InitConfig()
	balong.InitRouter()
	//wxmp.InitRouter()
	//wxmp.Init()
	//wxmp.InitSync()
	//数据库
	conf.InitDB()

	// beego.Router("/wxverify", &controllers.WxVerifyController{})
	// beego.AutoRouter(&controllers.WxVerifyController{})

	// beego.Router("/wx", &controllers.WxController{})
	// beego.Router("/wx/app/:appid", &controllers.WxController{}, "*:App")
	// beego.AutoRouter(&controllers.WxController{})

	beego.Router("/test", &controllers.TestController{})
	beego.AutoRouter(&controllers.TestController{})

	beego.Router("/suyuan", &controllers.SuYuanController{})
	beego.AutoRouter(&controllers.SuYuanController{})

	beego.Run()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	var srvname = beego.AppConfig.String("appname")

	svcConfig := &service.Config{
		Name:        srvname, //服务显示名称
		DisplayName: srvname, //服务名称
		Description: srvname, //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		beego.Error(err.Error())
	}

	if err != nil {
		beego.Error(err.Error())
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			fmt.Println("qtms服务安装成功")
			beego.AppConfig.Set("service", "1")
			beego.AppConfig.SaveConfigFile(beego.AppPath + "/conf/app.conf")
			return
		}

		if os.Args[1] == "uninstall" {
			s.Uninstall()
			fmt.Println("qtms服务卸载成功")
			beego.AppConfig.Set("service", "0")
			beego.AppConfig.SaveConfigFile(beego.AppPath + "/conf/app.conf")
			return
		}
	}

	err = s.Run()
	if err != nil {
		beego.Error(err)
	}
}
