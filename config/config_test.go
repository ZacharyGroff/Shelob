package config

import (
	"strings"
	"testing"
	"time"
)

func TestParseSeedPath(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := "seedPathTest"
	actual := config.SeedPath
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseSeedBuffer(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := 42
	actual := config.SeedBuffer
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestParseSleepSeconds(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := 43
	actual := config.SleepSeconds
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestParseInformSeconds(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := time.Duration(44)
	actual := config.InformSeconds
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestParseFlushToFile(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := false
	actual := config.FlushToFile
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestParseUrlFilterKeyword(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := "urlFilterTest"
	actual := config.UrlFilterKeyword
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
