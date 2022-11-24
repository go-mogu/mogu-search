package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/mgu-search/global"
	"github.com/go-mogu/mgu-search/internal/model"
	"github.com/go-mogu/mgu-search/internal/service"
	"github.com/go-mogu/mgu-search/pkg/consts/MessageConf"
	"github.com/go-mogu/mgu-search/pkg/response"
)

type searchController struct {
}

var SearchController = &searchController{}

// SearchBlog 博客搜索
// @Summary 博客搜索
// @Tags 索引相关
// @Description 博客搜索
// @Accept application/json
// @Produce application/json
// @Param     param  query      model.SearchBlog    true  "查询参数"
// @Success   200   {object}  response.JsonResponse{data=[]model.DocIndex}  "查询结果"
// @Router /search/elasticSearchBlog [get]
func (c *searchController) SearchBlog(context context.Context, ctx *app.RequestContext) {
	var param model.SearchBlog
	err := ctx.BindAndValidate(&param)
	if err != nil {
		response.BadRequestException(ctx, "")
	}
	blog, err := service.Search(global.Cfg.Search.Type).SearchBlog(param)
	if err != nil {
		response.FailedJson(ctx, err.Error(), nil)
	}
	response.SuccessJson(ctx, "", blog)
}

// DeleteElasticSearchByUidStr 批量删除博客索引
// @Summary 批量删除博客索引
// @Tags 索引相关
// @Description 批量删除博客索引
// @Accept application/json
// @Produce application/json
// @Param     uid  query      string    true  "删除uid"
// @Success   200   {object}  response.JsonResponse{msg=string}  "删除提示"
// @Router /search/deleteElasticSearchByUids [post]
func (c *searchController) DeleteElasticSearchByUidStr(context context.Context, ctx *app.RequestContext) {
	err := service.Search(global.Cfg.Search.Type).DeleteElasticSearchByUidStr(ctx.Query("uid"))
	if err != nil {
		response.FailedJson(ctx, err.Error(), nil)
	}
	response.SuccessJson(ctx, MessageConf.DELETE_SUCCESS, nil)
}

// DeleteElasticSearchByUid 删除博客索引
// @Summary 删除博客索引
// @Tags 索引相关
// @Description 删除博客索引
// @Accept application/json
// @Produce application/json
// @Param     uid  query      string    true  "删除uid"
// @Success   200   {object}  response.JsonResponse{msg=string}  "删除提示"
// @Router /search/deleteElasticSearchByUid [post]
func (c *searchController) DeleteElasticSearchByUid(context context.Context, ctx *app.RequestContext) {
	err := service.Search(global.Cfg.Search.Type).DeleteElasticSearchByUidStr(ctx.Query("uid"))
	if err != nil {
		response.FailedJson(ctx, err.Error(), nil)
	}
	response.SuccessJson(ctx, MessageConf.DELETE_SUCCESS, nil)
}

// AddElasticSearchIndexByUid 通过博客Uid添加索引
// @Summary 通过博客Uid添加索引
// @Tags 索引相关
// @Description 通过博客Uid添加索引
// @Accept application/json
// @Produce application/json
// @Param     uid  query      string    true  "博客uid"
// @Success   200   {object}  response.JsonResponse{msg=string}  "添加提示"
// @Router /search/addElasticSearchIndexByUid [post]
func (c *searchController) AddElasticSearchIndexByUid(context context.Context, ctx *app.RequestContext) {
	err := service.Search(global.Cfg.Search.Type).AddElasticSearchIndexByUid(context, ctx.Query("uid"))
	if err != nil {
		response.FailedJson(ctx, err.Error(), nil)
	}
	response.SuccessJson(ctx, MessageConf.DELETE_SUCCESS, nil)
}

// InitElasticSearchIndex 初始化索引
// @Summary 初始化索引
// @Tags 索引相关
// @Description 初始化索引
// @Accept application/json
// @Produce application/json
// @Success   200   {object}  response.JsonResponse{msg=string}  "初始化索引提示"
// @Router /search/initElasticSearchIndex [post]
func (c *searchController) InitElasticSearchIndex(context context.Context, ctx *app.RequestContext) {
	service.Search(global.Cfg.Search.Type).InitElasticSearchIndex(context, ctx)
}
