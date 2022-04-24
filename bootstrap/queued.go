package bootstrap

import (
	"chinadaily_com_cn/pkg/config"
	"chinadaily_com_cn/pkg/queued"
	"github.com/gocolly/colly/v2/queue"
	"sync"
)

var onceQueue sync.Once

// SetupQueued 初始化消息队列
func SetupQueued() {
	if Storage == nil {
		SetupRedisStorage()
	}
	onceQueue.Do(func() {
		var err error
		queued.Queued, err = queue.New(config.GetInt("spider.queue_count"), Storage)
		if err != nil {
			panic(err)
		}

	})
}
