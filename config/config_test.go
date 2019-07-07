package config

import (
	"testing"
	"strings"
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
