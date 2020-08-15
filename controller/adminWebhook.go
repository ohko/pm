package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pm/model"
	"strconv"
	"time"

	"github.com/ohko/hst"
)

// AdminWebhookController Webhook管理控制器
type AdminWebhookController struct {
	controller
}

// List 列表
func (o *AdminWebhookController) List(ctx *hst.Context) {
	ps, err := webhooks.List()
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{"ps": ps}, "admin/webhook/list.html")
}

// Detail 查看
func (o *AdminWebhookController) Detail(ctx *hst.Context) {
	wid, _ := strconv.Atoi(ctx.R.FormValue("WID"))

	p, err := webhooks.Get(wid)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, p, "admin/webhook/detail.html")
}

// Delete 删除
func (o *AdminWebhookController) Delete(ctx *hst.Context) {
	wid, _ := strconv.Atoi(ctx.R.FormValue("WID"))

	if err := webhooks.Delete(wid); err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	http.Redirect(ctx.W, ctx.R, "/admin_webhook/list", http.StatusFound)
}

// Push 接收
func (o *AdminWebhookController) Push(ctx *hst.Context) {
	reqHeadersBytes, err := json.MarshalIndent(ctx.R.Header, "", "  ")
	if err != nil {
		ctx.Data(500, err.Error())
	}
	bs, err := ioutil.ReadAll(ctx.R.Body)
	if err != nil {
		ctx.Data(500, err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")

	p := &model.Webhook{
		Head: string(reqHeadersBytes),
		Body: string(bs),
		Time: time.Now().In(loc),
	}
	if err := webhooks.Save(p); err != nil {
		ctx.Data(500, err.Error())
	}

	ctx.Data(200, "ok")
}
