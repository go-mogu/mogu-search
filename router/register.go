package router

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	_ "github.com/go-mogu/mgu-search/docs"
	"github.com/go-mogu/mgu-search/global"
	_ "github.com/go-mogu/mgu-search/internal/logic"
	baseClient "github.com/go-mogu/mgu-search/pkg/client"
	"github.com/go-mogu/mgu-search/pkg/response"
	"github.com/go-mogu/mgu-search/pkg/util"
	"github.com/go-mogu/mgu-search/router/routes"
	"github.com/go-mogu/mogu-registry/nacos"
	"github.com/hertz-contrib/requestid"
	"github.com/hertz-contrib/swagger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	swaggerFiles "github.com/swaggo/files"
)

func Register(port string) *server.Hertz {
	//获取本机ip
	addr := util.GetIpAddr()
	//nacos服务发现客户端
	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &global.Cfg.Nacos.Client,
			ServerConfigs: global.Cfg.Nacos.Server,
		})
	if err != nil {
		panic(err)
	}
	if global.Cfg.Server.Port != "" {
		port = global.Cfg.Server.Port
	}
	addr = addr + ":" + port
	//注册服务
	r := nacos.NewNacosRegistry(nacosCli)
	h := server.New(
		server.WithHostPorts("0.0.0.0"+":"+port),
		server.WithRegistry(r, &registry.Info{
			ServiceName: global.Cfg.Server.Name,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      1,
			Tags:        global.Cfg.Nacos.Discovery.Metadata,
		}),
	)
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(response.RecoveryHandler)))

	url := swagger.URL(fmt.Sprintf("http://0.0.0.0:%s/swagger/doc.json", port)) // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	// header add X-Request-Id
	h.Use(requestid.New())
	// 404 not found
	h.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		path := ctx.Request.URI().Path()
		method := ctx.Request.Method()
		response.NotFoundException(ctx, fmt.Sprintf("%s %s not found", method, path))
	})
	searchGroup := h.Group("")
	routes.InitSearchGroup(searchGroup)
	err = baseClient.InitClient()
	if err != nil {
		panic(err)
	}
	return h
}
