package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sjzdlm/db"
)

type SuYuanController struct {
	beego.Controller
}

func (c *SuYuanController) Get() {
	c.Ctx.WriteString("-")
}
func (c *SuYuanController) Hello() {
	c.Ctx.WriteString("ok")
}
func (c *SuYuanController) Dishes() {
	var code = c.GetString("code")
	if code == "" {
		c.Ctx.WriteString("[]")
		return
	}
	var sql = `
	select * from vpublicity_dishes  where ucode='93eb5b9593' order by week,mealtimes_id
	`
	var xx = db.NewDb("qtms")
	var list = db.Query2(xx, sql)
	var rst = `
	<style>
        body{background:#fff;}
        .clearfix{*zoom:1}
        .clearfix:after{content:'\20';display:block;clear:both;height:0;overflow:hidden;}
      ul li{list-style:none;}
        .fl{float:left;}
        .fr{float:right;}
        .fullImg{width:100%;}
        .fixed{position:fixed;z-index:1000;}
        .black{background: rgba(0, 0, 0, 0.7);}
        .borderb{border-bottom:1px solid #eae9e9;}
        .redColor{color:red;}
      .main{max-width:640px;margin:0 auto;padding-top:50px;}
        
            .close{margin:0 auto;margin-top:15px;width:45px;display:block;}
            

            .xinxi{top:0;left:0;right:0;bottom:0;}
            .xxcon{padding-bottom:10px;border-radius:3px;background:#fff;width:90%;margin:0 auto;margin-top:10%;}
            .xxcon .title{background:#01bf67;padding:8px 0;font-size:19px;border-radius:3px 3px 0 0 ;color:#fff;text-align:center;font-weight:bold;}
            .xxcon .weui-cells {margin-top:0;margin-bottom:5px;}
            .xxcon .weui-cells:after, .xxcon .weui-cells:before{border:none;}
            .xxcon .weui-label{color:#666666;width: 85px;}
            .returnT{width:100%;top:0;}
        .returnT_con{background:#fff;padding:10px;position:relative;}
        .returnT_con .renturn{position:absolute;left:10px;top:10px;}
        .returnT_con .renturn img{width:20px;}
        .returnT_con p{color:#494949;text-align:center;font-size:17px;letter-spacing:1px;text-overflow: ellipsis;white-space: nowrap;overflow: hidden;word-break: break-all; }

            .dpmc{position:relative;color:#fff;font-size:20px; letter-spacing:1px; font-weight:bold;}
            .dpmc img{display:block;}
            .dpmc p{position:absolute;left:5%;bottom:25%;}
            .cdgs_data{padding:10px;}
            .cdgs_data p{display:flex;align-items:center;color:#4e4e4e;}
            .cdgs_data p img{width:25px;margin-right:5px;}
        .caidan{background:#fcfcfc;padding:10px;}
        .cd_box{margin-bottom:10px;}
        .cd_box .title{position:relative;text-align:center;}
        .cd_box .title img{width:100%;}
        .cd_box .title span{position:absolute;left:11px;top:-4px;right:0;color:#4e4e4e;font-size:17px;}
        .cd_box ul li{float:left;width:48%;margin-bottom:10px;}
        .cd_box ul li:nth-child(2n){float:right;}
        .cd_box ul li img{width:100%;border-radius:5px;}
        .cd_box ul li p{color:#919191;font-size:16px;text-align:center;}

           
        </style>
	`
	var week = ""
	var meal = ""
	for i, row := range list {
		if row["weekname"] != week {
			if i > 0 {
				rst += `
						</ul>
					</div>
				</div>
				`
			}
			rst += `
			<div class="cdgs_data borderb">
                <p>
                    <img src="/images/suyuan/cdgs/data.png">
                    <span>` + row["weekname"] + `</span>
                </p>
            </div>
			`
			week = row["weekname"]
			meal = ""
		}
		if row["mealtimes_name"] != meal {
			if i > 0 {
				rst += `
						</ul>
					</div>
				</div>
				`
			}
			rst += `
			<div class="caidan">
                <div class="cd_box">
                    <div class="title">
                        <img src="/images/suyuan/cdgs/title1.png">
                        <span>` + row["mealtimes_name"] + `</span>
                    </div>
					<ul class="clearfix">`
		}
		rst += `
						<li>
                            <img src="` + row["foods_pic"] + `">
                            <p>` + row["foods_name"] + `</p>
                        </li>
		`
		meal = row["mealtimes_name"]

		// if row["mealtimes_name"] != meal {
		// 	rst += `
		//             </ul>
		//         </div>
		//     </div>
		// 	`
		// 	meal = row["mealtimes_name"]
		// }

	}
	c.Ctx.WriteString(rst)
}
