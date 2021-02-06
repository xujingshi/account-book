package http

import (
	"context"
	"go-common/library/ecode"
	binding "go-common/library/net/http/blademaster/binding"
	"net/http"
	"strconv"

	"account-book/internal/model"
	"account-book/internal/service"
	"go-common/library/conf/paladin"
	"go-common/library/log"
	"go-common/library/net/http/blademaster"
	bm "go-common/library/net/http/blademaster"
)

var svc *service.Service

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/v1/charge/detail")
	{
		g.POST("/add", add)
		g.GET("/list", list)
	}
}

func add(c *bm.Context) {
	req := &model.AddChargeDetailsReq{}
	if err := c.BindWith(req, binding.JSON, binding.Request); err != nil {
		return
	}
	mid, err := getMid(c)
	if err != nil {
		c.JSON(nil, ecode.Unauthorized)
		return
	}
	c.JSON(svc.AddChargeDetails(c, mid, req))
}

func list(c *bm.Context) {
	req := &model.GetDetailListReq{}
	if err := c.Bind(req); err != nil {
		return
	}
	mid, err := getMid(c)
	if err != nil {
		c.JSON(nil, ecode.Unauthorized)
		return
	}
	c.JSON(svc.GetDetailList(c, mid, req))
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func getCookieByKey(c context.Context, cookieName string) (cookieV string, err error) {
	ctx, ok := c.(*blademaster.Context)
	if ok {
		var cookie *http.Cookie
		cookie, err = ctx.Request.Cookie(cookieName)
		if err != nil {
			log.Error("getCookieByKey failed, expected cookie name=%+v, err=%+v", cookieName, err)
			return
		}
		cookieV = cookie.Value
	}
	return
}

func getMid(c context.Context) (mid int64, err error) {
	midStr, err := getCookieByKey(c, "mid")
	log.Infoc(c, "mid: %s, err: %v", midStr, err)
	if err != nil {
		return
	}
	return strconv.ParseInt(midStr, 10, 64)
}
