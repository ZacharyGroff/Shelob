package queue

import (
	"errors"
	"fmt"
	"net/url"
)

type SeedQueue struct {
	seeds chan url.URL
}

func NewSeedQueue(buffer int) *SeedQueue {
	seeds := make(chan url.URL, buffer)
	return &SeedQueue{seeds}
}

func (q SeedQueue) Get() (url.URL, error) {
	for {
		select {
		case url := <- q.seeds:
			return url, nil
		default:
			err := errors.New("No Urls in queue.")
			return url.URL{}, err
		}
	}
}

func (q SeedQueue) Put(url url.URL) error {
	select {
	case q.seeds <- url:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding url: %+v\n", url)
		return err
	}
}
