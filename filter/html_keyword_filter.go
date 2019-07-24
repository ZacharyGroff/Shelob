package filter

import (
	"strings"
	"github.com/ZacharyGroff/Shelob/config"
)

type HtmlKeywordFilter struct {
	Config *config.Config
}

func(h HtmlKeywordFilter) Filter(requestBody []byte) (bool, error) {
	keyword := h.Config.HtmlFilterKeyword
	return parseHtmlForKeyword(requestBody, keyword)
}

func parseHtmlForKeyword(requestBody []byte, keyword string) (bool, error) {
	if strings.Contains(string(requestBody), keyword) {
		return true, nil
	}
	return false, nil
}
