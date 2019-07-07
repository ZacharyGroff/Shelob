package queue

import (
	"testing"
	"net/url"
)

func TestPutSuccess(t *testing.T) {
	q := NewSeedQueue(1)
	url, _ := url.Parse("test.com/")
	err := q.Put(*url)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutError(t *testing.T) {
	q := NewSeedQueue(0)
	url, _ := url.Parse("test.com/")
	err := q.Put(*url)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetSuccess(t *testing.T) {
	expected, _ := url.Parse("test.com/")
	q := NewSeedQueue(1)
	q.Put(*expected)

	actual, _ := q.Get()
	if *expected != actual {
		t.Errorf("Expected: %+v\nActual: %+v\n", *expected, actual)
	}
}

func TestGetError(t *testing.T) {
	q := NewSeedQueue(0)
	url, _ := url.Parse("test.com/")
	q.Put(*url)

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}
