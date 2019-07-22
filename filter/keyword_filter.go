package filter

import (
	"bytes"
	"strings"
	"golang.org/x/net/html"
	"github.com/ZacharyGroff/Shelob/config"
)

type KeywordFilter struct {
	Config *config.Config
}

func (k KeywordFilter) Filter(requestBody []byte) (bool, error) {
	reader := bytes.NewReader(requestBody)
	tokenizer := html.NewTokenizer(reader)
	keyword := k.Config.FilterKeyword
	return parseForKeyword(tokenizer, keyword)
}

func parseForKeyword(t *html.Tokenizer, k string) (bool, error) {
	for {
		tokenType := t.Next()
		if tokenType == html.ErrorToken {
			return false, nil
		} else if tokenType == html.StartTagToken {
			token := t.Token()
			if containsKeyword(token, k) {
				return true, nil
			}
		}
	}
}

func containsKeyword(token html.Token, keyword string) bool {
	for _, attr := range token.Attr {
		if strings.Contains(attr.Val, keyword) {
			return true
		}
	}
	return false
}
