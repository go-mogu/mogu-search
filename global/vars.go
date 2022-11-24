package global

import (
	"context"
	"github.com/go-mogu/mgu-search/config"
	"github.com/spf13/viper"
	zinc "github.com/zinclabs/sdk-go-zincsearch"
)

var (
	Zinc     *zinc.APIClient
	ZincAuth context.Context
	Cfg      *config.Conf // yaml配置
	Viper    *viper.Viper
)
