package scheduler

import (
	"bufio"
	"strings"
	"testing"
	"net/url"
	"github.com/ZacharyGroff/Shelob/config"
	"github.com/ZacharyGroff/Shelob/queue"
)

func TestParseStringsForUrlsLength(t *testing.T) {
	expected := 2

	strings := []string{"https://test.com", "/test/path"}
	urls, _ := parseStringsForUrls(strings)
	actual := len(urls)

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)	
	} 
}

func TestParseStringsForUrlsCorrectUrl(t *testing.T) {
	expectedUrl, _ := url.Parse("https://test.com")
	expected := expectedUrl.String()

	strings := []string{"https://test.com"}
	urls, _ := parseStringsForUrls(strings)
	actual := urls[0].String()

	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)	
	} 
}

func TestParseStringsForUrlsSuccess(t *testing.T) {
	strings := []string{"https://test.com"}
	_, err := parseStringsForUrls(strings)

	if err != nil {
		t.Errorf("Expected nil but error returned.")
	} 
}

func TestParseStringsForUrlsError(t *testing.T) {
	strings := []string{"http\\:"}
	_, err := parseStringsForUrls(strings)

	if err == nil {
		t.Errorf("Expected error but nil returned.")
	} 
}

func TestGetLinesFromScannerLength(t *testing.T) {
	expected := 2

	reader := strings.NewReader("Test\nTest2")
	scanner := bufio.NewScanner(reader)
	lines := getLinesFromScanner(scanner)
	actual := len(lines)

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)	
	}
}

func TestGetLinesFromScannerCorrectLine(t *testing.T) {
	expected := "Test"

	reader := strings.NewReader("Test\nTest2")
	scanner := bufio.NewScanner(reader)
	lines := getLinesFromScanner(scanner)
	actual := lines[0]

	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)	
	}
}

func TestUpdateQueueLength(t *testing.T) {
	expected := 2
	
	url1, _ := url.Parse("https://test.com")
	url2, _ := url.Parse("test/path/")
	urls := []url.URL{*url1, *url2}

	config := config.Config{SeedBuffer: 5}
	queue := queue.NewSeedQueue(&config)

	scheduler := Scheduler{queue: queue}
	scheduler.updateQueue(urls)
	actual := scheduler.queue.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)	
	}
}

func TestUpdateQueueCorrectUrl(t *testing.T) {
	expectedUrl, _ := url.Parse("https://test.com")
	expected := expectedUrl.String()

	url1, _ := url.Parse("https://test.com")
	urls := []url.URL{*url1}

	config := config.Config{SeedBuffer: 5}
	queue := queue.NewSeedQueue(&config)

	scheduler := Scheduler{queue: queue}
	scheduler.updateQueue(urls)
	actualUrl, _ := scheduler.queue.Get()
	actual := actualUrl.String()

	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)	
	}
}
