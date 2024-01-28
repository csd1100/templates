package main

import (
	"log"

	"github.com/csd1100/templates/internal/parsers"
)

func main() {
	config, err := parsers.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	if config.Verbose {
		log.Printf("Config: %+v", config)
	}

	data, err := parsers.ParseTemplateFiles(config.ConfigFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	if config.Verbose {
		log.Printf("Data: %+v", data)
	}

	err = parsers.Generate(config, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
