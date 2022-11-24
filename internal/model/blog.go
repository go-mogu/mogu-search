package model

//type Blog struct {
//	Uid          string    `json:"uid"`
//	Title        string    `json:"title"`
//	Summary      string    `json:"summary"`
//	Content      string    `json:"content"`
//	TagUid       string    `json:"tagUid"`
//	ClickCount   int       `json:"clickCount"`
//	CollectCount int       `json:"collectCount"`
//	FileUid      string    `json:"fileUid"`
//	Status       int8      `json:"status"`
//	CreateTime   time.Time `json:"createTime"`
//	UpdatedAt    time.Time `json:"updateTime"`
//	AdminUid     string    `json:"adminUid"`
//	IsOriginal   string    `json:"isOriginal"`
//	Author       string    `json:"author"`
//	ArticlesPart string    `json:"articlesPart"`
//	BlogSortUid  string    `json:"blogSortUid"`
//	Level        int       `json:"level"`
//	IsPublish    string    `json:"isPublish"`
//	Sort         int       `json:"sort"`
//	OpenComment  string    `json:"openComment"`
//	Type         string    `json:"type"`
//	OutsideLink  string    `json:"outsideLink"`
//	Oid          int       `json:"oid"`
//	TagList      []Tag     `json:"tagList"`
//	PhotoList    []string  `json:"photoList"`
//	BlogSort     BlogSort  `json:"blogSort"`
//	BlogSortName string    `json:"blogSortName"`
//	PhotoUrl     string    `json:"photoUrl"`
//	ParseCount   int       `json:"parseCount"`
//	Copyright    string    `json:"copyright"`
//}

type Blog struct {
	Oid          int    `json:"oid"`
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	Content      string `json:"content"`
	TagUID       string `json:"tagUid"`
	BlogSortUID  string `json:"blogSortUid"`
	ClickCount   int    `json:"clickCount"`
	CollectCount int    `json:"collectCount"`
	FileUID      string `json:"fileUid"`
	AdminUID     string `json:"adminUid"`
	IsPublish    string `json:"isPublish"`
	IsOriginal   string `json:"isOriginal"`
	Author       string `json:"author"`
	ArticlesPart string `json:"articlesPart"`
	Level        int    `json:"level"`
	Sort         int    `json:"sort"`
	OpenComment  string `json:"openComment"`
	Type         string `json:"type"`
	TagList      []struct {
		Content    string `json:"content"`
		ClickCount int    `json:"clickCount"`
		Sort       int    `json:"sort"`
		UID        string `json:"uid"`
		Status     int    `json:"status"`
		CreateTime string `json:"createTime"`
		UpdateTime string `json:"updateTime"`
	} `json:"tagList"`
	PhotoList []string `json:"photoList"`
	BlogSort  struct {
		SortName   string `json:"sortName"`
		ClickCount int    `json:"clickCount"`
		Sort       int    `json:"sort"`
		UID        string `json:"uid"`
		Status     int    `json:"status"`
		CreateTime string `json:"createTime"`
		UpdateTime string `json:"updateTime"`
	} `json:"blogSort"`
	Copyright  string `json:"copyright"`
	UID        string `json:"uid"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type BlogPage[T any] struct {
	Data struct {
		Records          []T           `json:"records"`
		Total            int           `json:"total"`
		Size             int           `json:"size"`
		Current          int           `json:"current"`
		Orders           []interface{} `json:"orders"`
		OptimizeCountSQL bool          `json:"optimizeCountSql"`
		IsSearchCount    bool          `json:"isSearchCount"`
	} `json:"data"`
	Code string `json:"code"`
}

type ResultVo[T any] struct {
	Data T      `json:"data"`
	Code string `json:"code"`
}
