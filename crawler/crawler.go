package crawler

import (
	"log"
	"time"
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
	go crawler.Inform()
	crawler.scheduler.Start()
}

func (crawler Crawler) Stop() {
	log.Printf("Starting Shelob...\n")
}

func (crawler *Crawler) Inform() {
	for {
		log.Printf("Bytes downloaded: %d\n", *crawler.scheduler.BytesDownloaded)
		time.Sleep(crawler.config.InformSeconds * time.Second) 
	}
}
