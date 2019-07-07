package queue

import (
	"testing"
	"net/url"
	"github.com/ZacharyGroff/Shelob/config"
)

func TestPutSuccess(t *testing.T) {
	config := config.Config{"", 1}	
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")
	err := q.Put(*url)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutError(t *testing.T) {
	config := config.Config{"", 0}	
	q := NewSeedQueue(&config)
	url, _ := url.Parse("test.com/")
	err := q.Put(*url)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetSuccess(t *testing.T) {
	expected, _ := url.Parse("test.com/")
	config := config.Config{"", 1}	
	q := NewSeedQueue(&config)
	q.Put(*expected)

	actual, _ := q.Get()
	if *expected != actual {
		t.Errorf("Expected: %+v\nActual: %+v\n", *expected, actual)
	}
}

func TestGetError(t *testing.T) {
	config := config.Config{"", 0}	
	q := NewSeedQueue(&config)

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}
