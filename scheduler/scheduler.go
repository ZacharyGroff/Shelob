package scheduler

import (
	"log"
	"os"
	"bufio"
	"io"
	"fmt"
	"net/url"
	"net/http"
	"golang.org/x/net/html"
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
	err := scheduler.loadInitialSeeds()
	if err != nil {
		log.Fatalf("Failed to load initial seeds with error %s\n", err.Error())
	}
}

func (scheduler Scheduler) downloadAndParse(url url.URL) {
	reader, err := download(url)
	if err != nil {
		return
	}

	childUrls, err := parse(reader)
	if err != nil {
		return
	}

	for _, childUrl := range childUrls {
		scheduler.queue.Put(childUrl)
	}
}

func download(url url.URL) (io.Reader, error) {
	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func parse(reader io.Reader) ([]url.URL, error) {
	tokenizer := html.NewTokenizer(reader)
	urls, err := getUrls(tokenizer)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func getUrls(tokenizer *html.Tokenizer) ([]url.URL, error) {
	var urls []url.URL
	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()
			if isAnchor(token) {
				url, err := parseAnchorToken(token)
				if err != nil {
					return nil, err
				}
				urls = append(urls, url)
			}
		case tokenType == html.ErrorToken:
			return urls, nil
		}
	}
	err := fmt.Errorf("Reached end of tokenenizer without ErrorToken. Tokenizer: %+v\n", tokenizer)

	return nil, err
}

func isAnchor(token html.Token) bool {
	return token.Data == "a"
}

func parseAnchorToken(token html.Token) (url.URL, error) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			u, err := url.Parse(attr.Val)
			if err != nil {
				return url.URL{}, err
			}
			return *u, nil
		}
	}
	
	err := fmt.Errorf("Attempted to parse anchor token with no href. Token: %+v\n", token)
	return url.URL{}, err
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
