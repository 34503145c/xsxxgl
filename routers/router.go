package routers

import (
	"github.com/astaxie/beego"
	"github.com/lhtzbj12/sdrms/controllers/admin"
	"github.com/lhtzbj12/sdrms/controllers/user"
)

func init() {

	admin :=
		beego.NewNamespace("/admin",
			// //课程路由
			// beego.NSRouter("/course/index", &admin.CourseController{}, "*:Index"),
			// beego.NSRouter("/course/datagrid", &admin.CourseController{}, "Get,Post:DataGrid"),
			// beego.NSRouter("/course/edit/?:id", &admin.CourseController{}, "Get,Post:Edit"),
			// beego.NSRouter("/course/delete", &admin.CourseController{}, "Post:Delete"),
			// beego.NSRouter("/course/updateseq", &admin.CourseController{}, "Post:UpdateSeq"),
			// beego.NSRouter("/course/uploadimage", &admin.CourseController{}, "Post:UploadImage"),

			// //用户角色路由
			// beego.NSRouter("/role/index", &admin.RoleController{}, "*:Index"),
			// beego.NSRouter("/role/datagrid", &admin.RoleController{}, "Get,Post:DataGrid"),
			// beego.NSRouter("/role/edit/?:id", &admin.RoleController{}, "Get,Post:Edit"),
			// beego.NSRouter("/role/delete", &admin.RoleController{}, "Post:Delete"),
			// beego.NSRouter("/role/datalist", &admin.RoleController{}, "Post:DataList"),
			// beego.NSRouter("/role/allocate", &admin.RoleController{}, "Post:Allocate"),
			// beego.NSRouter("/role/updateseq", &admin.RoleController{}, "Post:UpdateSeq"),

			//资源路由
			beego.NSRouter("/resource/index", &admin.ResourceController{}, "*:Index"),
			beego.NSRouter("/resource/treegrid", &admin.ResourceController{}, "POST:TreeGrid"),
			beego.NSRouter("/resource/edit/?:id", &admin.ResourceController{}, "Get,Post:Edit"),
			beego.NSRouter("/resource/parent", &admin.ResourceController{}, "Post:ParentTreeGrid"),
			beego.NSRouter("/resource/delete", &admin.ResourceController{}, "Post:Delete"),
			//快速修改顺序
			beego.NSRouter("/resource/updateseq", &admin.ResourceController{}, "Post:UpdateSeq"),

			//通用选择面板
			beego.NSRouter("/resource/select", &admin.ResourceController{}, "Get:Select"),
			//用户有权管理的菜单列表（包括区域）
			beego.NSRouter("/resource/usermenutree", &admin.ResourceController{}, "POST:UserMenuTree"),
			beego.NSRouter("/resource/checkurlfor", &admin.ResourceController{}, "POST:CheckUrlFor"),

			// //后台用户路由
			// beego.NSRouter("/backenduser/index", &admin.BackendUserController{}, "*:Index"),
			// beego.NSRouter("/backenduser/datagrid", &admin.BackendUserController{}, "POST:DataGrid"),
			// beego.NSRouter("/backenduser/edit/?:id", &admin.BackendUserController{}, "Get,Post:Edit"),
			// beego.NSRouter("/backenduser/delete", &admin.BackendUserController{}, "Post:Delete"),
			// //后台用户中心
			// beego.NSRouter("/usercenter/profile", &admin.UserCenterController{}, "Get:Profile"),
			// beego.NSRouter("/usercenter/basicinfosave", &admin.UserCenterController{}, "Post:BasicInfoSave"),
			// beego.NSRouter("/usercenter/uploadimage", &admin.UserCenterController{}, "Post:UploadImage"),
			// beego.NSRouter("/usercenter/passwordsave", &admin.UserCenterController{}, "Post:PasswordSave"),

			beego.NSRouter("/home/index", &admin.HomeController{}, "*:Index"),
			beego.NSRouter("/home/login", &admin.HomeController{}, "*:Login"),
			beego.NSRouter("/home/dologin", &admin.HomeController{}, "Post:DoLogin"),
			beego.NSRouter("/home/logout", &admin.HomeController{}, "*:Logout"),
			// beego.NSRouter("/home/datareset", &admin.HomeController{}, "Post:DataReset"),

			beego.NSRouter("/home/404", &admin.HomeController{}, "*:Page404"),
			beego.NSRouter("/home/error/?:error", &admin.HomeController{}, "*:Error"),

			beego.NSRouter("/", &admin.HomeController{}, "*:Index"),
		)
	beego.Router("/", &user.HomeController{}, "*:Index")
	user :=
		beego.NewNamespace("/user",
			beego.NSRouter("/", &user.HomeController{}, "*:Index"),
			beego.NSRouter("/home/index", &user.HomeController{}, "*:Index"),
			beego.NSRouter("/home/login", &user.HomeController{}, "*:Login"),
			beego.NSRouter("/home/dologin", &user.HomeController{}, "Post:DoLogin"),
			beego.NSRouter("/home/logout", &user.HomeController{}, "*:Logout"),
		)
	//注册 namespace
	beego.AddNamespace(admin, user)
}
