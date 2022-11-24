package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/go-mogu/mgu-search/internal/controller"
	"github.com/go-mogu/mgu-search/internal/middleware"
)

// InitSearchGroup 初始化搜索服务接
func InitSearchGroup(r *route.RouterGroup) {
	group := r.Group("/search", middleware.SearchMiddle)
	group.GET("/elasticSearchBlog", controller.SearchController.SearchBlog)
	group.POST("/deleteElasticSearchByUids", controller.SearchController.DeleteElasticSearchByUidStr)
	group.POST("/deleteElasticSearchByUid", controller.SearchController.DeleteElasticSearchByUid)
	group.POST("/addElasticSearchIndexByUid", controller.SearchController.AddElasticSearchIndexByUid)
	group.POST("/initElasticSearchIndex", controller.SearchController.InitElasticSearchIndex)
}
