package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sjzmp/lib"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/sjzdlm/db"
)

var DBOrigin *xorm.Engine

func InitDB() {
	DBOrigin = db.NewDb("origin")
	if DBOrigin != nil {
		fmt.Println("[ok]DBOrigin数据库连接成功...")
		//go timerMsg()
	} else {
		fmt.Println("[error]DBOrigin数据库连接失败...")
	}
}

/*
  自动发送消息
*/
func timerMsg() {
	timer := time.NewTicker(10 * time.Second)
	for _ = range timer.C {
		fmt.Print("#")
		//获取微信信息
		var wx = db.First2(DBOrigin, "select * from wx_wechat where id='10001'")
		//读取一条没有发送的消息
		var m = db.First2(DBOrigin, "select * from sch_warnning where is_msgsend=0 order by id desc limit 1 ")
		//把状态修改成发送中
		db.Exec2(DBOrigin, "update sch_warnning set is_msgsend=2 where id=?", m["id"])
		if len(m) > 0 {
			//开始读取需要发送的
			var list = db.Query2(DBOrigin, "select * from wx_user where appid='wx76b08b4af5fda7ec' and company_id=?", m["recv_unit"])

			for _, row := range list {
				var data = make(map[string]string)
				data["first"] = m["msg_tip"]
				data["keyword1"] = m["send_title"]
				data["keyword2"] = m["send_name"]
				data["keyword3"] = m["logtime"]
				data["keyword4"] = m["send_content"]
				data["remark"] = "校园食材--关注校园食品安全的溯源平台"
				var err = wxutil.SendTemplateMsg(wx["access_token"], "U4PHEpWTtm6RMamFk3eVxW0nlCXkerXj74Tyo813ybk", row["openid"], "", data)
				if err != nil {
					fmt.Println("发送失败:", err.Error())
					//c.Ctx.WriteString("0")
					return
				}
			}
			//把状态修改成已发送
			db.Exec2(DBOrigin, "update sch_warnning set is_msgsend=1 where id=?", m["id"])
		}
		//如果电话不为空,则发送短信
		// if m["recv_mobile"] != "" {
		// 	SendSMS(m["recv_mobile"], m["send_content"]+"【来自:"+m["send_name"]+"】")
		// }
		//根据单位信息发给管理员短信,如果没有管理员取第一个人员的手机号发送
		var list = db.Query("select mobile from adm_user where mobile !='' and mobile is not null and company_id=? and is_manager=1", m["recv_unit"])
		if len(list) > 0 {
			for _, v := range list {
				var mobile = v["mobile"]
				if len(mobile) == 11 {
					SendSMS(mobile, m["send_content"]+"【来自:"+m["send_name"]+"】")
				}
			}
		} else {
			var u = db.First("select mobile from adm_user where mobile !='' and mobile is not null and company_id=? ", m["recv_unit"])
			if len(u) > 0 {
				if len(u["mobile"]) == 11 {
					SendSMS(u["mobile"], m["send_content"]+"(来自:"+m["send_name"]+")")
				}
			}
		}

	}
}

func SendSMS(mobile, msg string) string {
	client := http.Client{}
	data := "action=send&account=600010&password=ySTDeJ&mobile=" + mobile + "&content=" + msg
	data = data + "&extno=10690600010&rt=json"
	//请求
	request, err := http.NewRequest("POST", "http://47.104.227.172:7862/sms", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	//设置Content-Type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	//存库
	sms := SMSResult{}
	err1 := json.Unmarshal(body, &sms)
	if err1 == nil {
		var sql = `insert into sch_msg_mt(status,balance,mid,mobile,result)values(?,?,?,?,?)`
		db.Exec2(DBOrigin, sql, sms.Status, sms.Balance, sms.List[0].Mid, sms.List[0].Mobile, sms.List[0].Result)
	}
	return string(body)
}

type SMSResult struct {
	Status  string `json:"status"`
	Balance int    `json:"balance"`
	List    []struct {
		Mid    string `json:"mid"`
		Mobile string `json:"mobile"`
		Result int    `json:"result"`
	} `json:"list"`
}
