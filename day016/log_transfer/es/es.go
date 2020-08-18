package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// 将日志数据写入Elasticsearch

func Init() (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://172.16.0.109:9200"))
	if err != nil {
		// Handler error
		panic(err)
	}
	fmt.Println("connect to es success")
	// 创建一条数据
	//	禹都宫紫菀
	p1 := Person{Name: "guan", Age: 18, Married: true}
	put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())
	if err != nil {
		// Hander error
		panic(err)
	}
	fmt.Printf("Index user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return
}
