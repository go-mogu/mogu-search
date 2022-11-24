package test

import (
	"fmt"
	"github.com/expectedsh/go-sonic/sonic"
	"testing"
)

func TestSonic(t *testing.T) {

	ingester, err := sonic.NewIngester("120.48.7.239", 31184, "mogu2018")
	if err != nil {
		panic(err)
	}

	// I will ignore all errors for demonstration purposes

	_ = ingester.BulkPush("movies", "general", 3, []sonic.IngestBulkRecord{
		{"id:6ab56b4kk3", "Star wars"},
		{"id:5hg67f8dg5", "Spider man"},
		{"id:1m2n3b4vf6", "Batman"},
		{"id:68d96h5h9d0", "This is another movie"},
	}, sonic.LangAutoDetect)

	search, err := sonic.NewSearch("120.48.7.239", 31184, "mogu2018")
	if err != nil {
		panic(err)
	}

	results, _ := search.Query("movies", "general", "man", 10, 0, sonic.LangAutoDetect)

	fmt.Println(results)
}
