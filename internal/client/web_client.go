package client

import (
	"context"
	"encoding/json"
	"github.com/go-mogu/mgu-search/internal/model"
	baseClient "github.com/go-mogu/mgu-search/pkg/client"
	"github.com/go-mogu/mgu-search/pkg/util/gconv"
	"net/url"
)

type webClient struct{}

var WebClient = &webClient{}

const (
	webName         = "mogu-web"
	getSearchModel  = "/search/getSearchModel"
	getBlogByUid    = "/content/getBlogByUid"
	getBlogBySearch = "/index/getBlogBySearch"
)

func (c *webClient) GetBlogByUid(ctx context.Context, uid string) (blog model.Blog, err error) {
	body, err := baseClient.Get(ctx, webName, getBlogByUid, url.Values{"uid": {uid}})
	if err != nil {
		return
	}
	var blogPage model.ResultVo[model.Blog]
	err = json.Unmarshal(body, &blogPage)
	if err != nil {
		return
	}
	return blogPage.Data, nil
}

func (c *webClient) GetBlogBySearch(ctx context.Context, currentPage, pageSize int) (result []model.Blog, err error) {
	values := url.Values{}
	values.Set("currentPage", gconv.String(currentPage))
	values.Set("pageSize", gconv.String(pageSize))
	body, err := baseClient.Get(ctx, webName, getBlogBySearch, values)
	if err != nil {
		return
	}
	var blogPage model.BlogPage[model.Blog]
	err = json.Unmarshal(body, &blogPage)
	if err != nil {
		return
	}
	return blogPage.Data.Records, nil
}

func (c *webClient) GetSearchModel(ctx context.Context) (result []byte, err error) {
	result, err = baseClient.Get(ctx, webName, getSearchModel, nil)
	return
}
