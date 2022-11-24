package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/mgu-search/internal/model"
)

type ISearch interface {
	SearchBlog(blog model.SearchBlog) (result map[string]interface{}, err error)
	DeleteElasticSearchByUidStr(uid string) (err error)
	AddElasticSearchIndexByUid(ctx context.Context, uid string) (err error)
	InitElasticSearchIndex(ctx context.Context, requestContext *app.RequestContext)
}

var abstractSearch = map[string]ISearch{}

func Search(key string) ISearch {
	if abstractSearch[key] == nil {
		panic(fmt.Sprintf("implement not found for interface ISearch[%s], forgot register?", key))
	}
	return abstractSearch[key]
}
func RegisterSearch(key string, i ISearch) {
	abstractSearch[key] = i
}
