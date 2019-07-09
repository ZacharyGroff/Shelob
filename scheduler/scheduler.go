package scheduler

import (
	"log"
	"os"
	"bufio"
	"net/url"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/queue"
)

type Scheduler struct {
	config *config.Config
	queue queue.Queue
}

func NewScheduler(c *config.Config, q queue.Queue) *Scheduler {
	return &Scheduler{c, q}
}

func (scheduler Scheduler) Start() {
	log.Printf("Loading initial seeds...\n")
	err := scheduler.LoadInitialSeeds()
	if err != nil {
		log.Fatalf("Failed to load initial seeds with error %s\n", err.Error())
	}
}

func (scheduler Scheduler) LoadInitialSeeds() error {
	lines, err := getFileLines(scheduler.config.SeedPath)
	if err != nil {
		return err
	}

	initialSeeds, err := parseFileLines(lines)
	if err != nil {
		return err
	}

	for _, seed := range initialSeeds {
		scheduler.queue.Put(seed)
	}

	return nil
}

func getFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func parseFileLines(lines []string) ([]url.URL, error) {
	var urls []url.URL
	for _, line := range lines {
		url, err := url.Parse(line)
		if err != nil {
			return nil, err
		}
		urls = append(urls, *url)
	}
	
	return urls, nil
}
