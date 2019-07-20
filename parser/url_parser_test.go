package parser

import (
	"fmt"
	"bytes"
	"strings"
	"testing"
	"net/url"
	"golang.org/x/net/html"
)

func TestParseCorrectUrl(t *testing.T) {
	expectedUrl, _ := url.Parse("https://test.com/test/path.html")
	expected := expectedUrl.String()

	htm := `<!DOCTYPE html>
	<html>
	<head>
    	<title></title>
	</head>
	<body>
    	body content
    	<p>more <a href="/test/path.html">content</a></p>
	</body>
	</html>`
	
	body := []byte(htm)
	parent, _ := url.Parse("https://test.com/")
	
	parser := UrlParser{}
	urls := parser.Parse(body, *parent)
	actual := urls[0].String()

	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

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

func TestFillInPartialLinksCorrectLink(t *testing.T) {
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

func TestGetUrlsLength(t *testing.T) {
	expected := 2
	
	htm := `<!DOCTYPE html>
	<html>
	<head>
    	<title></title>
	</head>
	<body>
    	body content
    	<p>more <a href="https://tester.com/">content</a></p>
    	<p>This <a href="/foo"><em>important</em> link <br> to
	</body>
	</html>`
	
	body := []byte(htm)
	reader := bytes.NewReader(body)
	tokenizer := html.NewTokenizer(reader)
	urls := getUrls(tokenizer)
	actual := len(urls)

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestGetUrlsCorrectUrl(t *testing.T) {
	expectedUrl, _ := url.Parse("https://test.com/test/path.html")
	expected := expectedUrl.String()

	htm := `<!DOCTYPE html>
	<html>
	<head>
    	<title></title>
	</head>
	<body>
    	body content
    	<p>more <a href="https://test.com/test/path.html">content</a></p>
	</body>
	</html>`
	
	body := []byte(htm)
	reader := bytes.NewReader(body)
	tokenizer := html.NewTokenizer(reader)
	urls := getUrls(tokenizer)
	actual := urls[0].String()

	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetUrlSuccess(t *testing.T) {
	htm := `<a href="https://test.com/test/path.html">`
	
	body := []byte(htm)
	reader := bytes.NewReader(body)
	tokenizer := html.NewTokenizer(reader)
	tokenizer.Next()
	token := tokenizer.Token()
	_, err := getUrl(token)

	if err != nil {
		fmt.Println(err)
		t.Errorf("Expected nil but error returned.")
	}
}

func TestGetUrlError(t *testing.T) {
	htm := `<body>test content</body>`
	
	body := []byte(htm)
	reader := bytes.NewReader(body)
	tokenizer := html.NewTokenizer(reader)
	tokenizer.Next()
	token := tokenizer.Token()
	_, err := getUrl(token)

	if err == nil {
		t.Errorf("Expected error but nil returned.")
	}
}
