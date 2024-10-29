// Code generated by hertz generator.

package main

import (
	handler "github.com/ahang7/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
}
