package controller

import (
	"net/http"
	"strconv"

	"pm/model"

	"github.com/ohko/hst"
)

// AdminProjectController 项目管理控制器
type AdminProjectController struct {
	controller
}

// List 项目列表
func (o *AdminProjectController) List(ctx *hst.Context) {
	ps, err := projects.List()
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"ps": ps}, "admin/project/list.html")
}

// Add 增加项目
func (o *AdminProjectController) Add(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, nil, "admin/project/add.html")
	}

	p := &model.Project{
		Name: ctx.R.FormValue("Name"),
		Desc: ctx.R.FormValue("Desc"),
		Git:  ctx.R.FormValue("Git"),
		URL:  ctx.R.FormValue("URL"),
	}
	if err := projects.Save(p); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	http.Redirect(ctx.W, ctx.R, "/admin_project/list", http.StatusFound)
}

// Edit 编辑项目
func (o *AdminProjectController) Edit(ctx *hst.Context) {
	pid, _ := strconv.Atoi(ctx.R.FormValue("PID"))

	p, err := projects.Get(pid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, p, "admin/project/edit.html")
	}

	p.Name = ctx.R.FormValue("Name")
	p.Desc = ctx.R.FormValue("Desc")
	p.Git = ctx.R.FormValue("Git")
	p.URL = ctx.R.FormValue("URL")
	if err := p.Save(p); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	http.Redirect(ctx.W, ctx.R, "/admin_project/list", http.StatusFound)
}

// Delete 删除项目
func (o *AdminProjectController) Delete(ctx *hst.Context) {
	pid, _ := strconv.Atoi(ctx.R.FormValue("PID"))

	ts, err := tasks.ListByPID(pid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	if len(ts) > 0 {
		o.renderAdminError(ctx, "当前项目还有任务，不能删除")
	}

	if err := projects.Delete(pid); err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	http.Redirect(ctx.W, ctx.R, "/admin_project/list", http.StatusFound)
}
