package initialize

import (
	"fmt"
	"github.com/hespecial/gin-mall/global"
	"github.com/olivere/elastic/v7"
)

func InitEs() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("https://%s:%d", global.Config.Es.Host, global.Config.Es.Port)),
		elastic.SetBasicAuth(global.Config.Es.Username, global.Config.Es.Password),
		elastic.SetSniff(global.Config.Es.Sniffer),
	)
	if err != nil {
		panic(fmt.Sprintf("Fatal error elastic NewClient: %v", err))
	}
	return client
}
