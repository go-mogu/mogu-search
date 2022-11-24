package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-mogu/mgu-search/global"
	"github.com/go-mogu/mgu-search/internal/client"
	"github.com/go-mogu/mgu-search/internal/consts"
	"github.com/go-mogu/mgu-search/internal/model"
	"github.com/go-mogu/mgu-search/internal/service"
	"github.com/go-mogu/mgu-search/internal/util"
	"github.com/go-mogu/mgu-search/pkg/consts/SysConf"
	"github.com/go-mogu/mgu-search/pkg/response"
	"github.com/go-mogu/mgu-search/pkg/util/gconv"
	zinc "github.com/zinclabs/sdk-go-zincsearch"
	"strings"
)

func init() {
	service.RegisterSearch("zinc", NewZinc())
}

// NewZinc returns the interface service.
func NewZinc() *ZincSearch {
	return &ZincSearch{}
}

type ZincSearch struct{}

func (s *ZincSearch) SearchBlog(blog model.SearchBlog) (result map[string]interface{}, err error) {
	if blog.CurrentPage-1 > 0 {
		blog.CurrentPage = blog.CurrentPage - 1
	} else {
		blog.CurrentPage = 0
	}
	blog.Keywords = blog.Keywords + "*"
	query := zinc.V1ZincQuery{
		From:       &blog.CurrentPage,
		MaxResults: &blog.PageSize,
		SearchType: &blog.SearchType,
		Query: &zinc.V1QueryParams{
			Term: &blog.Keywords},
	}
	query.SetHighlight(handlerHighlight())
	searchResult, _, err := global.Zinc.Search.SearchV1(global.ZincAuth, consts.BlogIndex).Query(query).Execute()
	if err != nil {
		return nil, err
	}
	total := searchResult.Hits.Total.Value
	totalPage := *total / blog.PageSize
	if *total > 0 && totalPage == 0 {
		totalPage = 1
	}
	hits := searchResult.Hits.Hits
	blogList := make([]model.DocIndex, 0)
	for _, hit := range hits {
		doc := model.DocIndex{}
		err = gconv.Scan(hit.Source, &doc)
		if err != nil {
			return nil, err
		}
		//设置高亮
		if hit.Highlight[SysConf.SUMMARY] != nil {
			doc.Summary = gconv.String(gconv.SliceAny(hit.Highlight[SysConf.SUMMARY])[0])
		}
		if hit.Highlight[SysConf.TITLE] != nil {
			doc.Title = gconv.String(gconv.SliceAny(hit.Highlight[SysConf.TITLE])[0])
		}
		blogList = append(blogList, doc)
	}
	result = map[string]interface{}{}
	result[SysConf.TOTAL] = total
	result[SysConf.TOTAL_PAGE] = totalPage
	result[SysConf.PAGE_SIZE] = blog.PageSize
	result[SysConf.CURRENT_PAGE] = blog.CurrentPage + 1
	result[SysConf.BLOG_LIST] = blogList
	return
}

func handlerHighlight() zinc.MetaHighlight {
	fields := map[string]zinc.MetaHighlight{}
	fields[SysConf.TITLE] = zinc.MetaHighlight{
		PreTags:  []string{},
		PostTags: []string{},
	}
	fields[SysConf.SUMMARY] = zinc.MetaHighlight{
		PreTags:  []string{},
		PostTags: []string{},
	}
	return zinc.MetaHighlight{
		Fields:            &fields,
		FragmentSize:      nil,
		NumberOfFragments: nil,
		PreTags:           []string{"<span style='color:red'>"},
		PostTags:          []string{"</span>"},
	}
}

func (s *ZincSearch) DeleteElasticSearchByUidStr(uid string) (err error) {
	uidList := strings.Split(uid, SysConf.FILE_SEGMENTATION)
	for _, id := range uidList {
		_, _, err = global.Zinc.Document.Delete(global.ZincAuth, consts.BlogIndex, id).Execute()
		if err != nil {
			hlog.Error(err)
			return err
		}
	}
	return nil
}

func (s *ZincSearch) AddElasticSearchIndexByUid(ctx context.Context, uid string) (err error) {
	blog, err := client.WebClient.GetBlogByUid(ctx, uid)
	if err != nil {
		hlog.Error(err)
		return err
	}
	blogIndex := util.BuildBlog(blog)
	records := make([]map[string]interface{}, 0)
	records = append(records, gconv.Map(blogIndex))
	index := consts.BlogIndex
	_, _, err = global.Zinc.Document.Bulkv2(global.ZincAuth).Query(zinc.MetaJSONIngest{
		Index:   &index,
		Records: records,
	}).Execute()
	return
}

func (s *ZincSearch) InitElasticSearchIndex(ctx context.Context, requestContext *app.RequestContext) {
	_, _, err := global.Zinc.Index.Delete(global.ZincAuth, consts.BlogIndex).Execute()
	if err != nil {
		hlog.Error(err)
	}

	indexName := consts.BlogIndex
	shardNum := int32(1)
	storageType := global.Cfg.Search.Zinc.StorageType
	metaIndex := zinc.MetaIndexSimple{
		Mappings:    nil,
		Name:        &indexName,
		ShardNum:    &shardNum,
		StorageType: &storageType,
	}

	metaIndex.Mappings = initMappings()
	_, _, err = global.Zinc.Index.Create(global.ZincAuth).Data(metaIndex).Execute()
	if err != nil {
		hlog.Error(err)
		response.FailedJson(requestContext, err.Error(), nil)
	}
	page := 1
	row := 15
	size := 0
	blogList := make([]model.DocIndex, 0)
	for {
		// 查询blog信息
		list, err := client.WebClient.GetBlogBySearch(ctx, page, row)
		if err != nil {
			hlog.Error(err)
			response.FailedJson(requestContext, err.Error(), nil)
		}
		//构建blog
		size = len(list)
		for _, blog := range list {
			blogIndex := util.BuildBlog(blog)
			blogList = append(blogList, blogIndex)
		}
		// 翻页
		if size != 15 {
			goto SAVE
		}
		page++
	}
SAVE:
	records := gconv.Maps(blogList)
	index := consts.BlogIndex
	_, _, err = global.Zinc.Document.Bulkv2(global.ZincAuth).Query(zinc.MetaJSONIngest{
		Index:   &index,
		Records: records,
	}).Execute()
	if err != nil {
		response.FailedJson(requestContext, err.Error(), nil)
	}
	response.SuccessJson(requestContext, "初始化成功", nil)
}

func initMappings() map[string]interface{} {
	//添加需要高亮的即可
	fields := []string{SysConf.TITLE, SysConf.SUMMARY, SysConf.CONTENT}
	mappings := map[string]interface{}{}
	properties := map[string]interface{}{}
	for _, field := range fields {
		properties[field] = map[string]interface{}{
			"type":          "text",
			"indexName":     true,
			"store":         true,
			"highlightable": true,
		}
	}
	mappings["properties"] = properties
	return mappings
}
