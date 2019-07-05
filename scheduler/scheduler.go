package scheduler

import (
	"log"
	"net/url"
	"github.com/ZacharyGroff/Shelob/config"
)

type Scheduler struct {
	seeds []url.URL
}

func NewScheduler(config *config.Config) *Scheduler {
	seeds := getInitialSeeds(config.SeedPath)
	return &Scheduler{seeds}
}

func (scheduler Scheduler) Schedule() {
	log.Printf("Scheduling...\n")
}

func getInitialSeeds(path string) []url.URL {
	var urls []url.URL
	u, _ := url.Parse("https://google.com/")
	urls = append(urls, *u)
	return urls
}
