package user

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(Students))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func StudentsTBName() string {
	return TableName("students")
}

// // ResourceTBName 获取 Resource 对应的表名称
// func ResourceTBName() string {
// 	return TableName("resource")
// }

// // RoleTBName 获取 Role 对应的表名称
// func RoleTBName() string {
// 	return TableName("role")
// }

// // RoleResourceRelTBName 角色与资源多对多关系表
// func RoleResourceRelTBName() string {
// 	return TableName("role_resource_rel")
// }

// // RoleBackendUserRelTBName 角色与用户多对多关系表
// func RoleBackendUserRelTBName() string {
// 	return TableName("role_backenduser_rel")
// }

// func CourseTBName() string {
// 	return TableName("course")
// }
