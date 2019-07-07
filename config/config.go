package config

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type Config struct {
	SeedPath string `json:"seedPath"`
	SeedBuffer int `json:"seedBuffer"`
}

func (config *Config) parseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &config)
	
	return err
}

func NewConfig() *Config {
	config := Config{}
	err := config.parseConfig("config/conf.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}
