package parser

import (
	"strings"
	"testing"
	"net/url"
	"golang.org/x/net/html"
)

func TestParseAnchorTokenSuccess(t *testing.T) {
	token := html.Token{
		html.StartTagToken,
		1,
		"a",
		[]html.Attribute{
			html.Attribute{"", "href", "test.com"},
		},
	}
	_, err := parseAnchorToken(token)
	if err != nil {
		t.Errorf("Expected nil but error returned.")
	}
}

func TestParseAnchorTokenFailure(t *testing.T) {
	token := html.Token{
		html.StartTagToken,
		1,
		"a",
		[]html.Attribute{
			html.Attribute{"", "", ""},
		},
	}
	_, err := parseAnchorToken(token)
	if err == nil {
		t.Errorf("Expected error but nil returned.")
	}
}

func TestFillInPartialLinks(t *testing.T) {
	expectedUrl, _ := url.Parse("https://parent.com/test/path")
	expected := expectedUrl.String()

	partialUrl, _ := url.Parse("/test/path")
	partialUrls := []url.URL{*partialUrl}
	parent, _ := url.Parse("https://parent.com/")
	completeUrls := fillInPartialLinks(partialUrls, *parent)
	actual := completeUrls[0].String()

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
