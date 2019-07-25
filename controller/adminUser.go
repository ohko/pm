package controller

import (
	"net/http"
	"strconv"

	"pm/model"
	"pm/util"

	"github.com/ohko/hst"
)

// AdminUserController 用户管理控制器
type AdminUserController struct {
	controller
}

// List 用户列表
func (o *AdminUserController) List(ctx *hst.Context) {
	us, err := users.List()
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"us": us}, "admin/user/list.html")
}

// Add 增加用户
func (o *AdminUserController) Add(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		us, err := users.List()
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}

		o.renderAdmin(ctx, map[string]interface{}{"us": us}, "admin/user/add.html")
	}

	puid, _ := strconv.Atoi(ctx.R.FormValue("ParentUID"))

	u := &model.User{
		User:      ctx.R.FormValue("User"),
		Name:      ctx.R.FormValue("Name"),
		Pass:      string(util.Hash([]byte(ctx.R.FormValue("Pass")))),
		Email:     ctx.R.FormValue("Email"),
		ParentUID: puid,
	}
	if err := users.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}

// Edit 编辑用户
func (o *AdminUserController) Edit(ctx *hst.Context) {
	uid, _ := strconv.Atoi(ctx.R.FormValue("UID"))
	puid, _ := strconv.Atoi(ctx.R.FormValue("ParentUID"))
	u, err := users.Get(uid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.R.Method == "GET" {
		us, err := users.List()
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}
		o.renderAdmin(ctx, map[string]interface{}{"us": us, "u": u}, "admin/user/edit.html")
	}

	pass := ctx.R.FormValue("Pass")
	if pass != "" {
		u.Pass = string(util.Hash([]byte(pass)))
	}
	u.User = ctx.R.FormValue("User")
	u.Name = ctx.R.FormValue("Name")
	u.Email = ctx.R.FormValue("Email")
	u.ParentUID = puid
	if err := u.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}

// Delete 删除用户
// func (o *AdminUserController) Delete(ctx *hst.Context) {
// 	uid, _ := strconv.Atoi(ctx.R.FormValue("UID"))
// 	if err := users.Delete(uid); err != nil {
// 		o.renderAdminError(ctx, err.Error())
// 	}

// 	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
// }
