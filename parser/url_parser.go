package parser

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"golang.org/x/net/html"	
	"github.com/ZacharyGroff/Shelob/config"
)

type UrlParser struct {
	Config *config.Config
}

func NewUrlParser(config *config.Config) *UrlParser {
	return &UrlParser{config}
}

func (urlParser UrlParser) Parse(b []byte, parent url.URL) []url.URL {
	reader := bytes.NewReader(b)
	tokenizer := html.NewTokenizer(reader)
	urls := getUrls(tokenizer)
	urls = fillInPartialLinks(urls, parent)

	return urls
}

func fillInPartialLinks(urls []url.URL, parent url.URL) []url.URL {
	var completeUrls []url.URL
	for _, url := range urls {
		if isMissingHostname(url) {
			url.Host = parent.Host
		}
		if isMissingScheme(url) {
			url.Scheme = parent.Scheme
		}
		completeUrls = append(completeUrls, url)
	}
	return completeUrls
}

func getUrls(tokenizer *html.Tokenizer) []url.URL {
	var urls []url.URL
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return urls
		} else if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			url, err := getUrl(token)
			if err == nil {
				urls = append(urls, url)
			}
		}
	}
}

func getUrl(token html.Token) (url.URL, error) {
	if isAnchor(token) {
		return parseAnchorToken(token)
	}
	return url.URL{}, errors.New("Unhandled tag token reached.")
}

func isMissingHostname(url url.URL) bool {
	return url.Host == ""
}

func isMissingScheme(url url.URL) bool {
	return url.Scheme == ""
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
