package filter

import (
	"testing"
	"github.com/ZacharyGroff/Shelob/config"
)

func TestUrlFilterSuccess(t *testing.T) {
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
	requestBody := []byte(htm)
	
	config := config.Config{UrlFilterKeyword: "test"}
	filter := UrlKeywordFilter{&config}
	_, err := filter.Filter(requestBody)
	
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestUrlFilterTrueHref(t *testing.T) {
	expected := true

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
	requestBody := []byte(htm)
	
	config := config.Config{UrlFilterKeyword: "test"}
	filter := UrlKeywordFilter{&config}
	actual, _ := filter.Filter(requestBody)

	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestUrlFilterFalseKeywordNotPresent(t *testing.T) {
	expected := false

	htm := `<!DOCTYPE html>
	<html>
	<head>
    	<title></title>
	</head>
	<body>
    	body content
    	<p>more <a href="https://notkeyword.com/">content</a></p>
    	<p>This <a href="/foo"><em>important</em> link <br> to
	</body>
	</html>`
	requestBody := []byte(htm)
	
	config := config.Config{UrlFilterKeyword: "test"}
	filter := UrlKeywordFilter{&config}
	actual, _ := filter.Filter(requestBody)

	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestUrlFilterFalseKeywordInContent(t *testing.T) {
	expected := false

	htm := `<!DOCTYPE html>
	<html>
	<head>
    	<title></title>
	</head>
	<body>
    	body content
    	<p>test</p>
    	<p>This <a href="/foo"><em>test</em> link <br> to
	</body>
	</html>`
	requestBody := []byte(htm)
	
	config := config.Config{UrlFilterKeyword: "test"}
	filter := UrlKeywordFilter{&config}
	actual, _ := filter.Filter(requestBody)

	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}
