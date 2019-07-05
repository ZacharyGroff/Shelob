// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/crawler"
	"github.com/ZacharyGroff/Shelob/scheduler"
)

// Injectors from wire.go:

func InitializeShelob() crawler.Crawler {
	configConfig := config.NewConfig()
	schedulerScheduler := scheduler.NewScheduler(configConfig)
	crawlerCrawler := crawler.NewCrawler(configConfig, schedulerScheduler)
	return crawlerCrawler
}
