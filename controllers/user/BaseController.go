package user

import (
	"fmt"
	"strings"

	"github.com/lhtzbj12/sdrms/enums"

	"github.com/lhtzbj12/sdrms/models/common"
	"github.com/lhtzbj12/sdrms/models/user"

	//	"github.com/lhtzbj12/sdrms/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string        //当前控制名称
	actionName     string        //当前action名称
	curUser        user.Students //当前用户信息
}

func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	c.adapterUserInfo()
}

// checkLogin判断用户是否登录，未登录则跳转至登录页面
// 一定要在BaseController.Prepare()后执行
func (c *BaseController) checkLogin() {
	if c.curUser.Id == 0 {
		//登录页面地址
		urlstr := c.URLFor("user.HomeController.Login") + "?url="
		//登录成功后返回的址为当前
		returnURL := c.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if c.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := c.Ctx.Input.Refer()
			c.jsonResult(enums.JRCode302, "请登录", urlstr+returnURL)
		}
		c.Redirect(urlstr+returnURL, 302)
		c.StopRun()
	}
}

//从session里取用户信息
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession("students")
	if a != nil {
		c.curUser = a.(user.Students)
		c.Data["students"] = a
	}
}

//SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
func (c *BaseController) setStudent2Session(userId int) error {
	m, err := user.StudentsOne(userId)
	if err != nil {
		return err
	}
	//获取这个用户能获取到的所有资源列表
	// resourceList := user.ResourceTreeGridByUserId(userId, 1000)
	// for _, item := range resourceList {
	// 	m.ResourceUrlForList = append(m.ResourceUrlForList, strings.TrimSpace(item.UrlFor))
	// }
	c.SetSession("students", *m)
	return nil
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "user/layout/layout_page.html"
	switch {
	case len(template) == 1:
		tplName = "user/" + template[0]
	case len(template) == 2:
		tplName = "user/" + template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = "user/" + ctrlName + "/" + actionName + ".html"
	}
	c.Layout = layout
	c.TplName = tplName
}
func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &common.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) pageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *BaseController) pageLogin() {
	url := c.URLFor("HomeController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}
