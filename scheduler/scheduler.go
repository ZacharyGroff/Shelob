package scheduler

import (
	"log"
	"os"
	"bufio"
	"io"
	"time"
	"net/url"
	"net/http"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/queue"
	"github.com/ZacharyGroff/Shelob/parser"
)

type Scheduler struct {
	config *config.Config
	queue queue.Queue
}

func NewScheduler(c *config.Config, q *queue.SeedQueue) *Scheduler {
	return &Scheduler{c, q}
}

func (scheduler Scheduler) Start() {
	log.Printf("Loading initial seeds...\n")
	err := scheduler.loadInitialSeeds()
	if err != nil {
		log.Fatalf("Failed to load initial seeds with error %s\n", err.Error())
	}

	scheduler.Crawl()
}

func (scheduler Scheduler) Crawl() {
	urlParser := parser.NewParser(scheduler.config)
	for {
		seed, err := scheduler.queue.Get()
		if err != nil {
			scheduler.sleep()
			continue
		}
		reader, err := download(seed)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		childUrls := urlParser.Parse(reader, seed)
		scheduler.updateQueue(childUrls)
	}
}

func (scheduler Scheduler) sleep() {
	seconds := scheduler.config.SleepSeconds
	log.Printf("No urls in queue... sleeping for: %d seconds.\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
}

func (scheduler Scheduler) updateQueue(urls []url.URL) {
	for _, url := range urls { 
		err := scheduler.queue.Put(url)
		if err != nil {
			log.Println("Queue full. Flushing...")
			scheduler.queue.Flush()
			scheduler.queue.Put(url)
		}
	}
}

func download(url url.URL) (io.Reader, error) {
	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (scheduler Scheduler) loadInitialSeeds() error {
	lines, err := getFileLines(scheduler.config.SeedPath)
	if err != nil {
		return err
	}

	initialSeeds, err := parseFileLines(lines)
	if err != nil {
		return err
	}

	scheduler.updateQueue(initialSeeds)
	
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
