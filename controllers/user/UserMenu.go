package user

type UserMenu struct {
	Id       int8
	MenuName string
	Pid      int64
	//MenuType int8
	Url      string
	Children []UserMenu
}

// type TreeList struct {
// 	Id       int8
// 	MenuName string
// 	Pid      int64
// 	//MenuType int8
// 	Url      string
// 	Children []*TreeList `json:"children"`
// }

func GetuserMenu() []UserMenu {
	TreeList := make([]UserMenu, 0)
	child := make([]UserMenu, 0)
	memu := UserMenu{
		Id:       1,
		MenuName: "基本信息",
		Pid:      0,
		Url:      "/user/baseinfo",
		//Children: make([]UserMenu, 0),
	}
	TreeList = append(TreeList, memu)
	memu = UserMenu{
		Id:       1,
		MenuName: "课程信息",
		Pid:      0,
		Url:      "/user/baseinfo",
	}
	child = make([]UserMenu, 0)
	child = append(child, UserMenu{
		Id:       1,
		MenuName: "课程信息1",
		Pid:      0,
		Url:      "/user",
	})
	memu.Children = child
	TreeList = append(TreeList, memu)

	memu = UserMenu{
		Id:       1,
		MenuName: "成绩信息",
		Pid:      0,
		Url:      "/user/baseinfo",
		Children: make([]UserMenu, 0),
	}
	child = make([]UserMenu, 0)
	child = append(child, UserMenu{
		Id:       1,
		MenuName: "成绩信息1",
		Pid:      0,
		Url:      "/user",
	})
	memu.Children = child
	TreeList = append(TreeList, memu)
	return TreeList
}

// userMenu := {
// 	   Id:'1'
//        MenuName: '基本信息'
//        Pid:      '0'
//        Url:  '/user/baseinfo'
//        Children: '',
//       	Id:'2'
//        MenuName: '课程信息'
//        Pid:      '0'
//        Url:  '/user/kecheng'
//        Children: '',
//     Id:'3'
//        MenuName: '成绩信息'
//        Pid:      '0'
//        Url:  '/user/source'
//        Children: '',
// }

// return userMenu
