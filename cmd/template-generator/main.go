package main

import (
	"fmt"
	"log"

	"github.com/csd1100/templates/internal/parsers"
)

func main() {
	config, err := parsers.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("\n Config: %+v\n", config)

	data, err := parsers.ParseTemplateFiles(config.ConfigFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("\n Data: %+v\n", data)

	err = parsers.Generate(config, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
