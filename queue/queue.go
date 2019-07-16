package queue

import (
	"net/url"
)

type Queue interface {
	Size() int
	Get() (url.URL, error)
	Put(url.URL) error
	Flush() error 
}
