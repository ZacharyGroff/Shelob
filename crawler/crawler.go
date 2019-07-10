package crawler

import (
	"log"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/scheduler"
)

type Crawler struct {
	config *config.Config
	scheduler *scheduler.Scheduler
}

func NewCrawler(config *config.Config, scheduler *scheduler.Scheduler) Crawler {
	return Crawler{config, scheduler}
}

func (crawler Crawler) Start() {
	log.Printf("Starting Shelob...\n")
	crawler.scheduler.Start()
}

func (crawler Crawler) Stop() {
	log.Printf("Starting Shelob...\n")

}
