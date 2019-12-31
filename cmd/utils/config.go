package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Bitcoin struct {
		Host string
		User string
		Pass string
	}
	Ethereum struct {
		Host            string
		ContractAddress string
	}
}

func Load(filename string) (Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, err
	}
	return config, err
}
