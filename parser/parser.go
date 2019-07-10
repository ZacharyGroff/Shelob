package parser

import (
	"io"
	"fmt"
	"net/url"
	"golang.org/x/net/html"	
)

type Parser struct {
	config *config.Config
}

func Parse(reader io.Reader) ([]url.URL, error) {
	tokenizer := html.NewTokenizer(reader)
	urls, err := getUrls(tokenizer)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func getUrls(tokenizer *html.Tokenizer) ([]url.URL, error) {
	var urls []url.URL
	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()
			if isAnchor(token) {
				url, err := parseAnchorToken(token)
				if err != nil {
					return nil, err
				}
				urls = append(urls, url)
			}
		case tokenType == html.ErrorToken:
			return urls, nil
		}
	}
	err := fmt.Errorf("Reached end of tokenenizer without ErrorToken. Tokenizer: %+v\n", tokenizer)

	return nil, err
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
