package controller

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"pm/model"

	"github.com/ohko/hst"
	"github.com/ohko/logger"
)

var (
	ll          *logger.Logger
	sessionName = "TPLER"

	users    = model.NewUser()
	members  = model.NewMember()
	projects = model.NewProject()
	tasks    = model.NewTask()
	webhooks = model.NewWebhook()
)

type controller struct{}

// 渲染错误页面
func (o *controller) renderAdminError(ctx *hst.Context, data interface{}) {
	ctx.HTML2(200, "layout/admin.html", data, "admin/error.html")
}

// 渲染成功页面
func (o *controller) renderAdminSuccess(ctx *hst.Context, data interface{}) {
	ctx.HTML2(200, "layout/admin.html", data, "admin/success.html")
}

// 渲染后台模版
func (o *controller) renderAdmin(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/admin.html", data, names...)
}

// 渲染前台模版
func (o *controller) renderDefault(ctx *hst.Context, data interface{}, names ...string) {
	ctx.HTML2(200, "layout/default.html", data, names...)
}

// Start 启动WEB服务
func Start(addr, sessionPath, oauth2Server string, lll *logger.Logger) {
	ll = lll
	oauth2Init(oauth2Server)
	oauthStateString = time.Now().Format("20060102150405")

	// hst对象
	s := hst.New(nil)

	// 禁止显示Route日志
	// s.DisableRouteLog = true
	s.SetLogger(ioutil.Discard)

	// HTML模版
	s.SetDelims("{[{", "}]}")
	s.SetTemplatePath("./view/")

	// favicon.ico
	s.Favicon()

	// Session
	// s.SetSession(hst.NewSessionMemory())
	s.SetSession(hst.NewSessionFile("", "/", sessionName, sessionPath, time.Hour*24*30))

	// 静态文件
	s.StaticGzip("/public/", "./public/")

	// 注册自动路由
	s.RegisterHandle(
		[]hst.HandlerFunc{checkAdminLogined},
		&IndexController{},
		&AdminController{},
		&AdminUserController{},
		&AdminProjectController{},
		&AdminTaskController{},
		&AdminWebhookController{},
		&Oauth2Controller{},
	)

	// 设置模版函数
	s.SetTemplateFunc(map[string]interface{}{
		"json": func(x interface{}) string {
			bs, err := json.Marshal(x)
			if err != nil {
				return err.Error()
			}
			return string(bs)
		},
		"show_time": func(x interface{}) string {
			t, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 -0700", (x.(time.Time)).String())
			return t.Format("2006-01-02 15:04:05")
		},
		"show_commits_url": func(x string) string {
			var de map[string]interface{}
			if err := json.Unmarshal([]byte(x), &de); err != nil {
				return err.Error()
			}
			if v, ok := de["commits"]; ok {
				return (((v.([]interface{}))[0]).(map[string]interface{}))["url"].(string)
			}
			return ""
		},
		"show_commits_message": func(x string) template.HTML {
			var de map[string]interface{}
			if err := json.Unmarshal([]byte(x), &de); err != nil {
				return template.HTML(err.Error())
			}
			if v, ok := de["commits"]; ok {
				return template.HTML(strings.ReplaceAll((((v.([]interface{}))[0]).(map[string]interface{}))["message"].(string), "\n", "<br>"))
			}
			return ""
		},
		"show_pusher": func(x string) string {
			var de map[string]interface{}
			if err := json.Unmarshal([]byte(x), &de); err != nil {
				return err.Error()
			}
			if v, ok := de["pusher"]; ok {
				return (v.(map[string]interface{}))["username"].(string)
			}
			return ""
		},
	})

	// 启动web服务
	go s.ListenHTTP(addr)

	// 优雅关闭
	hst.Shutdown(time.Second*5, s)
}

// 登录检查
func checkAdminLogined(ctx *hst.Context) {

	if u, err := url.ParseRequestURI(ctx.R.RequestURI); err == nil {
		// 排除路径
		for _, v := range []string{
			"/",
			"/admin/login",
			"/admin_webhook/push",
			"/oauth2/login",
			"/oauth2/callback",
		} {
			if u.Path == v {
				return
			}
		}
	}

	if v, err := ctx.SessionGet("Member"); err == nil && v != nil {
		return
	}

	if strings.Contains(ctx.R.Header.Get("Accept"), "application/json") {
		ctx.JSON2(200, -1, "Please login")
	} else {
		uri := ctx.R.Host + ctx.R.RequestURI
		if ctx.R.TLS == nil {
			uri = "http://" + uri
		} else {
			uri = "https://" + uri
		}
		http.Redirect(ctx.W, ctx.R, "/admin/login?callback="+url.QueryEscape(uri), 302)
		ctx.Close()
	}
}
