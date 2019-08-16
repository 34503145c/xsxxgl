package user

import (
	"strings"

	"github.com/lhtzbj12/sdrms/enums"
	"github.com/lhtzbj12/sdrms/models/user"
	"github.com/lhtzbj12/sdrms/utils"
	//"github.com/lhtzbj12/sdrms/controllers/base"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()
	c.Data["name"] = "beego.me"

	c.setTpl()
}
func (c *HomeController) Page404() {
	c.setTpl()
}
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}
func (c *HomeController) Login() {

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "user/home/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "user/home/login_footerjs.html"
	c.setTpl("home/login.html", "shared/layout_base.html")
}
func (c *HomeController) DoLogin() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	userpwd = utils.String2md5(userpwd)
	user, err := user.StudentOneByStuNumber(username, userpwd)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			c.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setStudent2Session(user.Id)
		//获取用户信息
		c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "用户名或者密码错误", userpwd)
	}
}

func (c *HomeController) Logout() {
	user := user.Students{}
	c.SetSession("students", user)
	c.pageLogin()
}

// func (c *HomeController) DataReset() {
// 	if ok, err := user.DataReset(); ok {
// 		c.jsonResult(enums.JRCodeSucc, "初始化成功", "")
// 	} else {
// 		c.jsonResult(enums.JRCodeFailed, "初始化失败,可能原因:"+err.Error(), "")
// 	}

// }
