package config

import (
	"sync"
	"io/ioutil"
	"encoding/json"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	YoutubeApiKey string `json:"youtube_api_key"`
	SecretHMACKey string `json:"secret_hmac_key"`
	ApiPort       int    `json:"api_port"`
}

func Get() *Config {
	once.Do(func() {
		config = &Config{}

		file, err := ioutil.ReadFile("config/config.json")
		if err != nil {
			panic(err.Error())
		}

		if err := json.Unmarshal(file, config); err != nil {
			panic(err.Error())
		}
	})

	return config
}
