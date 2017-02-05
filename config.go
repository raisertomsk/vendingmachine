package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	Listen string `json:"listen"`
}

func NewConfig(path *string) (*Config, error) {
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, errors.New("Error open config file: " + err.Error())
	}
	config := new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, errors.New("Malformed config file: " + err.Error())
	}
	return config, nil
}
