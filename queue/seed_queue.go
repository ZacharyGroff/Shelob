package queue

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"net/url"
	"github.com/ZacharyGroff/Shelob/config"
)

type SeedQueue struct {
	seeds chan url.URL
	config *config.Config
}

func NewSeedQueue(config *config.Config) *SeedQueue {
	seeds := make(chan url.URL, config.SeedBuffer)
	return &SeedQueue{seeds, config}
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

func (q SeedQueue) Flush() error {
	file, err := os.OpenFile(q.config.SeedPath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	initialSize := len(q.seeds)
	for i := 0; i < initialSize; i++ {
		url, err := q.Get()
		if err != nil {
			return err
		}

		fmt.Fprintln(writer, url) 
	}
	
	return writer.Flush()
}
