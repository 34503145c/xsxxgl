package user

import (
	"github.com/astaxie/beego/orm"
	"github.com/lhtzbj12/sdrms/models/common"
)

// TableName 设置Students表名
func (a *Students) TableName() string {
	return StudentsTBName()
}

// BackendUserQueryParam 用于查询的类
type StudentsQueryParam struct {
	common.BaseQueryParam
	UserNameLike string //模糊查询
	RealNameLike string //模糊查询
	Mobile       string //精确查询
	SearchStatus string //为空不查询，有值精确查询
}

//  Students 实体类
type Students struct {
	//StudentId int    `orm:"column(student_id);pk"`
	Id int `orm:"column(student_id);pk"`
	//Id        int
	StuNumber string `orm:"size(32)"`
	Name      string `orm:"size(24)"`
	UserPwd   string `json:"-"`
	// IsSuper            bool
	Status int
	// Mobile             string                `orm:"size(16)"`
	// Email              string                `orm:"size(256)"`
	// Avatar             string                `orm:"size(256)"`
	//RoleIds            []int                 `orm:"-" form:"RoleIds"`
	//	RoleBackendUserRel []*RoleBackendUserRel `orm:"reverse(many)"` // 设置一对多的反向关系
	//	ResourceUrlForList []string              `orm:"-"`
	//CreateCourses      []*Course             `rom:"reverse(many)"` // 设置一对多的反向关系
}

func StudentsOne(id int) (*Students, error) {
	o := orm.NewOrm()
	m := Students{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// BackendUserOneByUserName 根据用户名密码获取单条
func StudentOneByStuNumber(username, userpwd string) (*Students, error) {
	m := Students{}
	err := orm.NewOrm().QueryTable(StudentsTBName()).Filter("stu_number", username).Filter("userpwd", userpwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
