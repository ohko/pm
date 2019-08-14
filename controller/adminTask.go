package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"pm/model"

	"github.com/ohko/hst"
)

// AdminTaskController 任务管理控制器
type AdminTaskController struct {
	controller
}

// List 任务列表
func (o *AdminTaskController) List(ctx *hst.Context) {
	ts, err := tasks.List()
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"ts": ts}, "admin/task/list.html")
}

// ListByProject 项目任务列表
func (o *AdminTaskController) ListByProject(ctx *hst.Context) {
	pid, _ := strconv.Atoi(ctx.R.FormValue("ProjectPID"))
	ts, err := tasks.ListByPID(pid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"from": "project", "ts": ts, "ProjectPID": pid}, "admin/task/list.html")
}

// ListByUser 成员任务列表
func (o *AdminTaskController) ListByUser(ctx *hst.Context) {
	uid, _ := strconv.Atoi(ctx.R.FormValue("UserUID"))
	ts, err := tasks.ListByUser(uid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"from": "user", "ts": ts, "UserUID": uid}, "admin/task/list.html")
}

// Add 增加任务
func (o *AdminTaskController) Add(ctx *hst.Context) {
	from := ctx.R.FormValue("from")
	pid, _ := strconv.Atoi(ctx.R.FormValue("ProjectPID"))
	uid, _ := strconv.Atoi(ctx.R.FormValue("UserUID"))
	progress, _ := strconv.Atoi(ctx.R.FormValue("Progress"))

	if ctx.R.Method == "GET" {

		ps, err := projects.List()
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}

		us, err := users.List()
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}

		o.renderAdmin(ctx, map[string]interface{}{"from": from, "ps": ps, "us": us, "ProjectPID": pid, "UserUID": uid}, "admin/task/add.html")
	}

	date := strings.Split(ctx.R.FormValue("Date"), " - ")
	t := &model.Task{
		ProjectPID: pid,
		UserUID:    uid,
		Name:       ctx.R.FormValue("Name"),
		Desc:       ctx.R.FormValue("Desc"),
		Git:        ctx.R.FormValue("Git"),
		Start:      date[0],
		End:        date[1],
		Progress:   progress,
	}
	if err := tasks.Save(t); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	switch from {
	case "project":
		http.Redirect(ctx.W, ctx.R, "/admin_task/list_by_project?ProjectPID="+ctx.R.FormValue("ProjectPID"), http.StatusFound)
	case "user":
		http.Redirect(ctx.W, ctx.R, "/admin_task/list_by_user?UserUID="+ctx.R.FormValue("UserUID"), http.StatusFound)
	default:
		http.Redirect(ctx.W, ctx.R, "/admin_task/list", http.StatusFound)
	}
}

// Edit 编辑任务
func (o *AdminTaskController) Edit(ctx *hst.Context) {
	from := ctx.R.FormValue("from")
	pid, _ := strconv.Atoi(ctx.R.FormValue("ProjectPID"))
	uid, _ := strconv.Atoi(ctx.R.FormValue("UserUID"))
	tid, _ := strconv.Atoi(ctx.R.FormValue("TID"))
	progress, _ := strconv.Atoi(ctx.R.FormValue("Progress"))
	date := strings.Split(ctx.R.FormValue("Date"), " - ")

	t, err := tasks.Get(tid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.R.Method == "GET" {

		ps, err := projects.List()
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}

		us, err := users.List()
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}

		o.renderAdmin(ctx, map[string]interface{}{"from": from, "ps": ps, "us": us, "t": t}, "admin/task/edit.html")
	}

	t.Name = ctx.R.FormValue("Name")
	t.Desc = ctx.R.FormValue("Desc")
	t.Git = ctx.R.FormValue("Git")
	t.ProjectPID = pid
	t.UserUID = uid
	t.Start = date[0]
	t.End = date[1]
	t.Progress = progress
	if err := t.Save(t); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	switch from {
	case "project":
		http.Redirect(ctx.W, ctx.R, fmt.Sprintf("/admin_task/list_by_project?ProjectPID=%d", t.ProjectPID), http.StatusFound)
	case "user":
		http.Redirect(ctx.W, ctx.R, "/admin_task/list_by_user?UserUID="+ctx.R.FormValue("UserUID"), http.StatusFound)
	default:
		http.Redirect(ctx.W, ctx.R, "/admin_task/list", http.StatusFound)
	}
}

// Delete 删除任务
func (o *AdminTaskController) Delete(ctx *hst.Context) {
	from := ctx.R.FormValue("from")
	tid, _ := strconv.Atoi(ctx.R.FormValue("TID"))
	if err := tasks.Delete(tid); err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	switch from {
	case "project":
		http.Redirect(ctx.W, ctx.R, "/admin_task/list_by_project?ProjectPID="+ctx.R.FormValue("ProjectPID"), http.StatusFound)
	case "user":
		http.Redirect(ctx.W, ctx.R, "/admin_task/list_by_user?UserUID="+ctx.R.FormValue("UserUID"), http.StatusFound)
	default:
		http.Redirect(ctx.W, ctx.R, "/admin_task/list", http.StatusFound)
	}
}
