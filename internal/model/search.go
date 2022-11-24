package model

type SearchBlog struct {
	Keywords    string `json:"keywords" query:"keywords,required"`
	SearchType  string `json:"searchType" default:"querystring"`
	CurrentPage int32  `json:"currentPage" default:"1"`
	PageSize    int32  `json:"pageSize" default:"10"`
}
