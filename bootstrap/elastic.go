package bootstrap

import (
	"chinadaily_com_cn/pkg/config"
	pkgelastic "chinadaily_com_cn/pkg/elastic"
	"github.com/olivere/elastic/v7"
	"time"
)

// SetupElastic 初始化 Elastic.
func SetupElastic() {
	pkgelastic.Options = []elastic.ClientOptionFunc{
		elastic.SetURL(config.GetString("elastic.host")),
		elastic.SetBasicAuth(
			config.GetString("elastic.username"),
			config.GetString("elastic.password"),
		),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(5 * time.Second),
	}
}
