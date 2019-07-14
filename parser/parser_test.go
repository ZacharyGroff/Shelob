package parser

import (
	"testing"
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
