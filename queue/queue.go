package queue

import (
	"net/url"
)

type Queue interface {
	Get() (url.URL, error)
	Put(url.URL) error
	Flush() error 
}
