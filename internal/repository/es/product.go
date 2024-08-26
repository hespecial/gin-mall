package es

import (
	"context"
	"encoding/json"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
	"github.com/olivere/elastic/v7"
)

func SearchProduct(keyword string, page, size int) (products []*model.Product, count int64, err error) {
	// 创建查询
	query := elastic.NewMatchQuery("title", keyword)

	// 执行搜索操作
	searchResult, err := global.EsClient.Search().
		Index("products"). // 搜索的索引
		Query(query). // 设定查询条件
		From(size * (page - 1)).
		Size(size).
		Do(context.Background()) // 执行查询
	if err != nil {
		return nil, 0, err
	}

	// 解析搜索结果
	for _, hit := range searchResult.Hits.Hits {
		var product *model.Product
		err = json.Unmarshal(hit.Source, &product)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	return products, searchResult.Hits.TotalHits.Value, nil
}
