package queue

import (
	"os"
	"testing"
	"net/url"
	"github.com/ZacharyGroff/Shelob/config"
)

func TestPutSuccess(t *testing.T) {
	config := config.Config{"", 1, 0}	
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")
	err := q.Put(*url)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutError(t *testing.T) {
	config := config.Config{"", 0, 0}	
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")
	err := q.Put(*url)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetSuccess(t *testing.T) {
	expected, _ := url.Parse("test.com/")
	config := config.Config{"", 1, 0}	
	q := NewSeedQueue(&config)
	q.Put(*expected)

	actual, _ := q.Get()
	if *expected != actual {
		t.Errorf("Expected: %+v\nActual: %+v\n", *expected, actual)
	}
}

func TestGetError(t *testing.T) {
	config := config.Config{"", 0, 0}
	q := NewSeedQueue(&config)

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestFlushSize(t *testing.T) {
	testPath := "seed_test.txt"
	os.Create(testPath)

	config := config.Config{testPath, 1, 0}
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")

	q.Put(*url)
	q.Flush()
	
	expected := 0
	actual := len(q.seeds)
	if expected != actual {
		os.Remove(testPath)
		t.Errorf("Expected: %d\tActual: %d\n", expected, actual)	
	}
	
}

func TestFlushSuccess(t *testing.T) {
	testPath := "seed_test.txt"
	f, err := os.Create(testPath)
	f.Close()

	config := config.Config{testPath, 1, 0}
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")

	q.Put(*url)
	err = q.Flush()
	
	if err != nil {
		os.Remove(testPath)
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testPath)
}

func TestFlushError(t *testing.T) {
	testPath := "seed_test.txt"

	config := config.Config{testPath, 1, 0}
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")

	q.Put(*url)
	err := q.Flush()
	
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestSizeZero(t *testing.T) {
	expected := 0

	config := config.Config{SeedBuffer: 5}
	q := NewSeedQueue(&config)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestSizeNotZero(t *testing.T) {
	expected := 2

	config := config.Config{SeedBuffer: 5}
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")

	q.Put(*url)
	q.Put(*url)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
