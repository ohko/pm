package model

import "fmt"

// Menu 菜单定义
type Menu struct {
	Class string // 图标类名
	Text  string // 文字
	Href  string // 链接
	Child []Menu `json:",omitempty"` // 子菜单
}

// GetAdminMenu ...
func (o *Menu) GetAdminMenu(who string) []Menu {
	projects := NewProject()
	ps, _ := projects.List()
	var psChild []Menu
	for _, v := range ps {
		psChild = append(psChild, Menu{Class: "fa-circle-o", Text: v.Name, Href: fmt.Sprintf("/admin_task/list_by_project?ProjectPID=%d", v.PID)})
	}

	users := NewUser()
	us, _ := users.List()
	var usChild []Menu
	for _, v := range us {
		usChild = append(usChild, Menu{Class: "fa-circle-o", Text: v.Name, Href: fmt.Sprintf("/admin_task/list_by_user?UserUID=%d", v.UID)})
	}

	return []Menu{
		Menu{Class: "fa-home", Text: "仪表盘", Href: "/admin/"},
		Menu{Class: "fa-cube", Text: "项目管理", Href: "/admin_project/list"},
		Menu{Class: "fa-cubes", Text: "项目任务列表", Child: psChild},
		Menu{Class: "fa-user", Text: "成员管理", Href: "/admin_user/list"},
		Menu{Class: "fa-users", Text: "成员任务列表", Child: usChild},
		Menu{Class: "fa-tasks", Text: "任务管理", Href: "/admin_task/list"},
		Menu{Class: "fa-lock", Text: "修改密码", Href: "/admin/password"},
		Menu{Class: "fa-share", Text: "退出:" + who, Href: "javascript:vueMenu.logout()"},
	}
}
