package util

import (
	"github.com/go-mogu/mgu-search/internal/model"
	"github.com/go-mogu/mgu-search/pkg/util/empty"
)

func BuildBlog(blog model.Blog) (blogIndex model.DocIndex) {
	//构建blog对象
	blogIndex = model.DocIndex{
		IndexId:     blog.UID,
		Id:          blog.UID,
		Uid:         blog.UID,
		Oid:         blog.Oid,
		Type:        blog.Type,
		Title:       blog.Title,
		Summary:     blog.Summary,
		Content:     blog.Content,
		BlogSortUid: blog.BlogSortUID,
		IsPublish:   blog.IsPublish,
		Author:      blog.Author,
		CreateTime:  blog.CreateTime,
		UpdateTime:  blog.UpdateTime,
	}

	if !empty.IsEmpty(blog.BlogSort) {
		blogIndex.BlogSortName = blog.BlogSort.SortName

	}
	if len(blog.TagList) > 0 {
		tagUidList := make([]string, 0)
		tagNameList := make([]string, 0)
		for _, tag := range blog.TagList {
			tagUidList = append(tagUidList, tag.UID)
			tagNameList = append(tagNameList, tag.Content)
		}
		blogIndex.TagUidList = tagUidList
		blogIndex.TagNameList = tagNameList
	}

	return
}
