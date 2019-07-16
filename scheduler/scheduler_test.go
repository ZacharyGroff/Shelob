package scheduler

import (
	"testing"
	"net/url"
)

func TestParseStringsForUrlsLength(t *testing.T) {
	expected := 2

	strings := []string{"https://test.com", "/test/path"}
	urls, _ := parseStringsForUrls(strings)
	actual := len(urls)

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)	
	} 
}

func TestParseStringsForUrlsCorrectUrl(t *testing.T) {
	expectedUrl, _ := url.Parse("https://test.com")
	expected := expectedUrl.String()

	strings := []string{"https://test.com"}
	urls, _ := parseStringsForUrls(strings)
	actual := urls[0].String()

	if expected != actual {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)	
	} 
}

func TestParseStringsForUrlsSuccess(t *testing.T) {
	strings := []string{"https://test.com"}
	_, err := parseStringsForUrls(strings)

	if err != nil {
		t.Errorf("Expected nil but error returned.")
	} 
}

func TestParseStringsForUrlsError(t *testing.T) {
	strings := []string{"http\\:"}
	_, err := parseStringsForUrls(strings)

	if err == nil {
		t.Errorf("Expected error but nil returned.")
	} 
}
