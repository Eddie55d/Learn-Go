package config

import (
	"encoding/json"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Database struct {
		Host           string `json:"host"`
		Port           string `json:"port"`
		Name           string `json:"name"`
		User           string `json:"user"`
		Password       string `json:"password"`
		SslMode        string `json:"sslmode"`
		ConnectTimeout string `json:"connect-timeout"`
	} `json:"database"`
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	return config, err
}
