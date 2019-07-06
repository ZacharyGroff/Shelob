//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/scheduler"
	"github.com/ZacharyGroff/Shelob/crawler"
	"github.com/ZacharyGroff/Shelob/queue"
)

func InitializeShelob() crawler.Crawler {
	wire.Build(crawler.NewCrawler, scheduler.NewScheduler, queue.NewSeedQueue,
		config.NewConfig)

	return crawler.Crawler{}
}
