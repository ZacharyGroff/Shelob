package scheduler

import (
	"log"
	"os"
	"bufio"
	"io"
	"time"
	"io/ioutil"
	"net/url"
	"net/http"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/queue"
	"github.com/ZacharyGroff/Shelob/parser"
)

type Scheduler struct {
	config *config.Config
	queue queue.Queue
	bytesDownloaded uint64
}

func NewScheduler(c *config.Config, q *queue.SeedQueue) *Scheduler {
	return &Scheduler{c, q, 0}
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
		go scheduler.incrementBytesDownloaded(reader)
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

func (scheduler Scheduler) loadInitialSeeds() error {
	lines, err := getFileLines(scheduler.config.SeedPath)
	if err != nil {
		return err
	}

	initialSeeds, err := parseStringsForUrls(lines)
	if err != nil {
		return err
	}

	scheduler.updateQueue(initialSeeds)
	
	return nil
}

func (scheduler Scheduler) incrementBytesDownloaded(r io.Reader) error {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	bytesInReader := uint64(len(bytes))
	scheduler.bytesDownloaded += bytesInReader

	return nil
}

func download(url url.URL) (io.Reader, error) {
	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func getFileLines(path string) ([]string, error) {
	scanner, err := getScannerFromFile(path)
	if err != nil {
		return nil, err
	}
	lines := getLinesFromScanner(scanner)

	return lines, nil
}

func getScannerFromFile(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return &bufio.Scanner{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	return scanner, nil
}

func getLinesFromScanner(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseStringsForUrls(lines []string) ([]url.URL, error) {
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
