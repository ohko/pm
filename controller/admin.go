package controller

import (
	"net/http"
	"strings"

	"pm/model"
	"pm/util"

	"github.com/ohko/hst"
)

// AdminController 默认管理控制器
type AdminController struct {
	controller
}

// Login 登录
func (o *AdminController) Login(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		ctx.HTML2(200, "layout/empty.html", map[string]string{
			"callback": ctx.R.FormValue("callback"),
		}, "admin/login.html")
	}

	ctx.R.ParseForm()
	user := ctx.R.FormValue("User")
	pass := ctx.R.FormValue("Password")
	callback := ctx.R.FormValue("callback")

	if err := members.Check(user, pass); err != nil {
		ctx.JSON2(200, 1, err.Error())
	}

	ctx.SessionSet("Member", user)

	if strings.Contains(ctx.R.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		ctx.JSON2(200, 0, "ok")
	} else if callback != "" {
		http.Redirect(ctx.W, ctx.R, callback, 302)
	} else {
		http.Redirect(ctx.W, ctx.R, "/admin/", 302)
	}
}

// Logout 登出
func (o *AdminController) Logout(ctx *hst.Context) {
	ctx.SessionDestory()
	http.Redirect(ctx.W, ctx.R, "/admin/login", 302)
	ctx.Close()
}

// Index ...
func (o *AdminController) Index(ctx *hst.Context) {
	m, _ := ctx.SessionGet("Member")
	o.renderAdmin(ctx, &model.Member{
		User: m.(string),
	}, "admin/index.html")
}

// Password ...
func (o *AdminController) Password(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, nil, "admin/password.html")
	}

	newPass := ctx.R.FormValue("Pass")
	if len(newPass) == 0 {
		o.renderAdminError(ctx, "新密码不能为空")
	}

	m, err := ctx.SessionGet("Member")
	if err != nil || m.(string) == "" {
		o.renderAdminError(ctx, "发现错误")
	}
	u := &model.Member{
		User: m.(string),
		Pass: string(util.Hash([]byte(newPass))),
	}
	if err := members.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	o.renderAdminSuccess(ctx, "密码修改成功")
}

// GetMenu 默认菜单
func (o *AdminController) GetMenu(ctx *hst.Context) {
	m, _ := ctx.SessionGet("Member")
	ctx.JSON2(200, 0, new(model.Menu).GetAdminMenu(m.(string)))
}

// Form ...
func (o *AdminController) Form(ctx *hst.Context) {
	o.renderAdmin(ctx, nil, "admin/form.html")
}

// Table ...
func (o *AdminController) Table(ctx *hst.Context) {
	o.renderAdmin(ctx, nil, "admin/table.html")
}
