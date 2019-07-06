package queue

import (
	"log"
	"time"
	"net/url"
)

type SeedQueue struct {
	seeds chan url.URL
}

func NewSeedQueue() *SeedQueue {
	seeds := make(chan url.URL, 1000)
	return &SeedQueue{seeds}
}

func (q SeedQueue) Get() url.URL {
	for {
		select {
		case url := <- q.seeds:
			return url
		default:
			log.Print("No Urls in queue. Sleeping...\n")
			time.Sleep(30 * time.Second)
		}
	}
}

func (q SeedQueue) Put(url url.URL) {
	select {
	case q.seeds <- url:
		return
	default:
		log.Printf("No room in buffer. Discarding url: %+v\n", url)
	}
}
