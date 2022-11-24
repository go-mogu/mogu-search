package client

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/go-mogu/mgu-search/global"
	"github.com/go-mogu/mgu-search/pkg/util/encoding/gurl"
	"github.com/go-mogu/mogu-registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"net/url"
)

var (
	scheme     = "http"
	BaseClient *client.Client
)

func InitClient() error {
	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &global.Cfg.Nacos.Client,
			ServerConfigs: global.Cfg.Nacos.Server,
		})
	if err != nil {
		panic(err)
	}
	r := nacos.NewNacosResolver(nacosCli)
	baseClient, err := client.NewClient()
	if err != nil {
		return err
	}
	baseClient.Use(sd.Discovery(r))
	BaseClient = baseClient
	return nil
}

func Get(ctx context.Context, serviceName, url string, data url.Values) (result []byte, err error) {
	//req := protocol.AcquireRequest()
	//resp := protocol.AcquireResponse()
	//req.SetOptions(config.WithSD(true))
	if data != nil {
		url = url + "?" + gurl.BuildQuery(data)
	}
	url = fmt.Sprintf("%s://%s%s", scheme, serviceName, url)
	//req.SetRequestURI(url)
	//req.ParseURI()
	_, result, err = BaseClient.Get(ctx, nil, url, config.WithSD(true))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return
	}
	return
}
