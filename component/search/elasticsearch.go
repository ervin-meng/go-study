package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

type account struct {
	Name string `json:"name"`
}

func main() {
	logger := log.New(os.Stdout, "elastic", log.LstdFlags)
	client, _ := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	////查询数据
	//query := elastic.NewMatchQuery("name", "bobby")
	//result, _ := client.Search().Index("account").Query(query).Do(context.Background())
	//
	//fmt.Println(result.Hits.TotalHits.Value)
	//
	//for _, value := range result.Hits.Hits {
	//	account := &account{}
	//	_ = json.Unmarshal(value.Source, account)
	//	fmt.Println(account.Name)
	//}
	////插入数据
	//account1 := account{
	//	Name: "ervin",
	//}
	//
	//result1,_ := client.Index().Index("account").BodyJson(account1).Do(context.Background())
	//
	//fmt.Println(result1.Index,result1.Id)
	////创建索引
	goodsMapping := `
    {
		"mappings":{
			"properties":{
				"name":{
					"type":"text",
					"analyzer":"ik_max_word"
				},
				"id":{
					"type":"integer"
				}
			}
		}
	}`

	createIndex, _ := client.CreateIndex("goods").BodyString(goodsMapping).Do(context.Background())

	if !createIndex.Acknowledged {

	}
}
