package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PORT        string `yaml:"PORT"`
	HOST        string `yaml:"HOST"`
	SERVER_TYPE string `yaml:"SERVER_TYPE"`
}

var config Config

func Configure(args []string) {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	yamlDecoder := yaml.NewDecoder(f)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) > 1 {
		config.HOST = args[1]
	}
}
