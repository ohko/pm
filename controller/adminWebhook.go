package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pm/model"
	"sort"
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
	tags := webhooks.GetTags()
	sort.Strings(tags)

	ps, err := webhooks.List(ctx.R.FormValue("tag"))
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	o.renderAdmin(ctx, map[string]interface{}{
		"tags": tags,
		"ps":   ps,
	}, "admin/webhook/list.html")
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

	pusher := ""
	var de map[string]interface{}
	if err := json.Unmarshal(bs, &de); err != nil {
		ctx.Data(500, err.Error())
	}
	if v, ok := de["pusher"]; ok {
		pusher = (v.(map[string]interface{}))["full_name"].(string)
		if pusher == "" {
			pusher = (v.(map[string]interface{}))["username"].(string)
		}
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")

	p := &model.Webhook{
		Head: string(reqHeadersBytes),
		Body: string(bs),
		Time: time.Now().In(loc),
		Tag:  pusher,
	}
	if err := webhooks.Save(p); err != nil {
		ctx.Data(500, err.Error())
	}

	webhooks.Clean(90)

	ctx.Data(200, "ok")
}

// Tags tags
func (o *AdminWebhookController) Tags(ctx *hst.Context) {
	tags := webhooks.GetTags()

	ctx.Data(200, tags)
}
