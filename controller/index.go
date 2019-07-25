package controller

import (
	"github.com/ohko/hst"
)

// IndexController 默认主页控制器
type IndexController struct {
	controller
}

// 渲染模版
func (o *IndexController) render(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/default.html", data, names...)
}

// Index 默认主页
func (o *IndexController) Index(ctx *hst.Context) {
	o.render(ctx, nil, "page/index.html")
}
