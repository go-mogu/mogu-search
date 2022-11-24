package test

import (
	"context"
	"fmt"
	"github.com/blugelabs/bluge"
	"log"
	"testing"
	"time"
)

func TestBluge(t *testing.T) {
	// write index
	writeIndex("./data/bluge/")
	// batch insert
	batch("./data/bluge/")
	// search
	search("./data/bluge/")
}

// 创建索引
func writeIndex(indexPath string) {
	config := bluge.DefaultConfig(indexPath)
	writer, err := bluge.OpenWriter(config)
	if err != nil {
		log.Fatalf("error opening writer: %v", err)
	}
	defer writer.Close()

	// 新建文档
	doc := bluge.NewDocument("example").
		AddField(bluge.NewTextField("name", "bluge")).AddField(bluge.NewDateTimeField("created_at", time.Now()))

	err = writer.Update(doc.ID(), doc)
	if err != nil {
		log.Fatalf("error updating document: %v", err)
	}
}

// 批量创建
func batch(indexPath string) {
	writer, err := bluge.OpenWriter(bluge.DefaultConfig(indexPath))
	batch := bluge.NewBatch()
	for i := 0; i < 10; i++ {
		doc := bluge.NewDocument(fmt.Sprintf("example_%d", i)).
			AddField(bluge.NewTextField(fmt.Sprintf("field_%d", i), fmt.Sprintf("value_%d", i%2))).AddField(bluge.NewDateTimeField("created_at", time.Now()))
		batch.Insert(doc)
	}
	err = writer.Batch(batch)
	if err != nil {
		log.Fatalf("error executing batch: %v", err)
	}
	batch.Reset()
}

// 查询
func search(indexPath string) {
	config := bluge.DefaultConfig(indexPath)
	reader, err := bluge.OpenReader(config)

	if err != nil {
		log.Fatalf("error getting index reader: %v", err)
	}
	defer reader.Close()

	query := bluge.NewMatchQuery("value_1").SetField("field_1")
	request := bluge.NewTopNSearch(10, query).
		WithStandardAggregations()
	documentMatchIterator, err := reader.Search(context.Background(), request)
	if err != nil {
		log.Fatalf("error executing search: %v", err)
	}
	match, err := documentMatchIterator.Next()
	for err == nil && match != nil {
		err = match.VisitStoredFields(func(field string, value []byte) bool {
			fmt.Printf("match: %s:%s\n", field, string(value))
			return true
		})
		if err != nil {
			log.Fatalf("error loading stored fields: %v", err)
		}
		fmt.Println(match)
		match, err = documentMatchIterator.Next()
	}
	if err != nil {
		log.Fatalf("error iterator document matches: %v", err)
	}
}
