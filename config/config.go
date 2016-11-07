package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var c *Config

// Config represents a config struct.
type Config struct {
	SL struct {
		APIKey        string   `json:"api_key"`
		TimeWindow    int      `json:"time_window"`
		TransportMode []string `json:"transport_mode"`
	}
	Pushover struct {
		AppKey  string `json:"app_key"`
		UserKey string `json:"user_key"`
	}
}

// Init will initialize the config file.
func Init(s string) {
	path := "config.json"

	if len(s) > 0 {
		path = s
	}

	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(fmt.Sprintf("Config error: %v\n", err))
		return
	}

	var config *Config

	json.Unmarshal(file, &config)

	c = config
}

// Get will return the config struct.
func Get() *Config {
	if c == nil {
		Init("")
	}

	if len(c.SL.TransportMode) == 0 {
		c.SL.TransportMode = []string{"train"}
	}

	return c
}
