package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Address   string
	Port      string
	Home      string
	Templates []string
	Apps      []string
	DB        DBConfig
}

type DBConfig struct {
	Type string
	User string
	Name string
	IP   string
	Pass string
}

func Load_config(location string) Config {
	file, _ := os.Open(location)
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config
}
