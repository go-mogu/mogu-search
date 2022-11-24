package model

type DocIndex struct {
	IndexId      string   `json:"_id"`
	Id           string   `json:"id"`
	Uid          string   `json:"uid"`
	Oid          int      `json:"oid"`
	Type         string   `json:"type"` //资源类型: 文章、动态、问答、面经
	Title        string   `json:"title"`
	Summary      string   `json:"summary"`
	Content      string   `json:"content"`
	SortName     string   `json:"sortName"`
	BlogSortName string   `json:"blogSortName"`
	BlogSortUid  string   `json:"blogSortUid"`
	IsPublish    string   `json:"isPublish"`
	CreateTime   string   `json:"createTime"`
	UpdateTime   string   `json:"updateTime"`
	Author       string   `json:"author"` //作者
	PhotoUrl     string   `json:"photoUrl"`
	TagNameList  []string `json:"tagNameList"` //标签名称
	TagUidList   []string `json:"tagUidList"`  //标签id
}
