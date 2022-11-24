package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/mgu-search/global"
	"github.com/go-mogu/mgu-search/internal/client"
	"github.com/go-mogu/mgu-search/pkg/consts/SysConf"
	"github.com/go-mogu/mgu-search/pkg/util/gconv"
)

func SearchMiddle(c context.Context, ctx *app.RequestContext) {

	result, err := client.WebClient.GetSearchModel(c)
	if err != nil {
		return
	}
	resultMap := gconv.MapStrStr(result)
	if SysConf.SUCCESS == resultMap[SysConf.CODE] {
		//暂时只支持zinc
		if resultMap[SysConf.DATA] == "1" {
			global.Cfg.Search.Type = "zinc"
			ctx.Next(c)
		}
		return
	}

}
