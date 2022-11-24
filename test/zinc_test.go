package test

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-mogu/mgu-search/bootstrap"
	"github.com/go-mogu/mgu-search/config"
	"github.com/go-mogu/mgu-search/global"
	"github.com/go-mogu/mgu-search/pkg/consts/SysConf"
	"github.com/go-mogu/mgu-search/pkg/util/gconv"
	"github.com/meilisearch/meilisearch-go"
	zinc "github.com/zinclabs/sdk-go-zincsearch"
	"io"
	"os"
	"testing"
)

func init() {
	config.ConfEnv = "dev"
	bootstrap.BootService()
}

func TestIndex(t *testing.T) {

	t.Run("create", func(t *testing.T) {
		indexStr := `{
    "name": "blog",
    "storage_type": "disk",
    "shard_num": 1,
    "mappings": {
        "properties": {
            "@timestamp": {
                "type": "date",
                "index": true,
                "store": false,
                "sortable": true,
                "aggregatable": true,
                "highlightable": false
            },
            "_id": {
                "type": "keyword"
            },
            "content": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "title": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            },
            "summary": {
                "type": "text",
                "index": true,
                "store": true,
                "highlightable": true
            }
        }
    }
}
`
		simple := zinc.MetaIndexSimple{}
		err := gconv.Scan(indexStr, &simple)
		if err != nil {
			panic(err)
		}
		index, resp, err := global.Zinc.Index.Create(global.ZincAuth).Data(simple).Execute()
		if err != nil {
			panic(err)
		}
		hlog.Infof("index: %v\n", index)
		hlog.Infof("resp: %v\n", resp)
	})
	t.Run("exists", func(t *testing.T) {
		result, resp, err := global.Zinc.Index.Exists(global.ZincAuth, "1article").Execute()
		if err != nil && err.Error() == "404 Not Found" {
			hlog.Error("index不存在")
		}
		hlog.Infof("result: %v\n", result)
		hlog.Infof("resp: %v\n", resp)
	})
	t.Run("createDoc", func(t *testing.T) {
		result, resp, err := global.Zinc.Document.Index(global.ZincAuth, "blog").Document(map[string]interface{}{
			"title":   "tttt",
			"content": "222222222",
			"status":  1,
		}).Execute()
		if err != nil {
			hlog.Error("index不存在")
		}
		hlog.Infof("result: %v\n", result)
		hlog.Infof("resp: %v\n", resp)
	})

	t.Run("search", func(t *testing.T) {
		queryMap := map[string]interface{}{
			"from":        0,
			"max_results": 20,
			"search_type": "querystring",
			"sort_fields": []string{"-title"},
			"query": map[string]interface{}{
				"term": "牛",
			},
		}
		fields := map[string]zinc.MetaHighlight{}
		fields[SysConf.TITLE] = zinc.MetaHighlight{
			PreTags:  []string{},
			PostTags: []string{},
		}
		fields[SysConf.CONTENT] = zinc.MetaHighlight{
			PreTags:  []string{},
			PostTags: []string{},
		}
		fields[SysConf.SUMMARY] = zinc.MetaHighlight{
			PreTags:  []string{},
			PostTags: []string{},
		}
		query := zinc.V1ZincQuery{}
		query.SetHighlight(zinc.MetaHighlight{
			Fields:            &fields,
			FragmentSize:      nil,
			NumberOfFragments: nil,
			PreTags:           []string{"<span style='color:red'>"},
			PostTags:          []string{"</span>"},
		})
		err := gconv.Scan(queryMap, &query)
		if err != nil {
			panic(err)
		}
		result, resp, err := global.Zinc.Search.SearchV1(global.ZincAuth, "blog").Query(query).Execute()
		if err != nil {
			panic(err)
		}
		hlog.Infof("result: %v\n", result)
		hlog.Infof("resp: %v\n", resp)
	})
}

func TestMeSearch(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "mogu2018",
	})
	search, err := client.Index("movies").Search("american ninja", &meilisearch.SearchRequest{})
	if err != nil {
		return
	}
	t.Logf("%v", search)

}

func TestDoc(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "mogu2018",
	})
	documents := []map[string]interface{}{
		{
			"id":           287947,
			"title":        "牛逼要死了",
			"poster":       "https://image.tmdb.org/t/p/w1280/xnopI5Xtky18MPhK40cZAGAOVeV.jpg",
			"overview":     "窝草，怎么这么强我屮艸芔茻",
			"release_date": "2019-03-23",
		},
	}
	addDocuments, err := client.Index("movies").AddDocuments(documents)
	if err != nil {
		return
	}
	t.Logf("%v", addDocuments)

}

func TestAddDoc(t *testing.T) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://search.ithhit.cn",
		APIKey: "mogu2018",
	})

	jsonFile, _ := os.Open("movies.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var movies []map[string]interface{}
	json.Unmarshal(byteValue, &movies)

	_, err := client.Index("movies").AddDocuments(movies)
	if err != nil {
		panic(err)
	}
}
