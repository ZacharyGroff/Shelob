package filter

import (
	"fmt"
	"bytes"
	"strings"
	"golang.org/x/net/html"
	"github.com/ZacharyGroff/Shelob/config"
)

type UrlKeywordFilter struct {
	Config *config.Config
}

func (k UrlKeywordFilter) Filter(requestBody []byte) (bool, error) {
	reader := bytes.NewReader(requestBody)
	tokenizer := html.NewTokenizer(reader)
	keyword := k.Config.UrlFilterKeyword
	return parseTokensForKeyword(tokenizer, keyword)
}

func parseTokensForKeyword(t *html.Tokenizer, k string) (bool, error) {
	for {
		tokenType := t.Next()
		if tokenType == html.ErrorToken {
			return false, nil
		} else if tokenType == html.StartTagToken {
			token := t.Token()
			if isAnchor(token) && containsKeyword(token, k) {
				return true, nil
			}
		}
	}
}

func isAnchor(token html.Token) bool {
	return token.Data == "a"
}

func containsKeyword(token html.Token, keyword string) bool {
	for _, attr := range token.Attr {
		fmt.Println(attr.Val)
		if strings.Contains(attr.Val, keyword) {
			return true
		}
	}
	return false
}
