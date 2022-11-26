package service

import (
	"context"
	"fmt"
	"github.com/go-mogu/mgu-search/internal/model"
)

type ISearch interface {
	SearchBlog(ctx context.Context, blog model.SearchBlog) (result map[string]interface{}, err error)
	DeleteElasticSearchByUidStr(ctx context.Context, uid string) (err error)
	AddElasticSearchIndexByUid(ctx context.Context, uid string) (err error)
	InitElasticSearchIndex(ctx context.Context)
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
