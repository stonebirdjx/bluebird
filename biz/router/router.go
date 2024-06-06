package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/stonebirdjx/bluebird/biz/handler"
)

// CustomizedRegister 注册路由信息
func CustomizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
}
