//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/scheduler"
	"github.com/ZacharyGroff/Shelob/crawler"
)

func InitializeShelob() crawler.Crawler {
	wire.Build(crawler.NewCrawler, scheduler.NewScheduler, config.NewConfig)

	return crawler.Crawler{}
}
