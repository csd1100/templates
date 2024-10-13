package parsers

import (
	"encoding/json"
	"fmt"
	"os"
)

type TemplateFile struct {
	Real         string            `json:"real"`
	Template     string            `json:"template"`
	Replacements map[string]string `json:"replacements"`
}

type TemplateFiles struct {
	Files []TemplateFile `json:"template-files"`
}

func ParseTemplateFiles(templateFile string) (*TemplateFiles, error) {
	var templateFiles TemplateFiles
	contents, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file '%v'", templateFile)
	}

	err = json.Unmarshal(contents, &templateFiles)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config '%v', due to error: '%w'", templateFile, err)
	}

	return &templateFiles, nil
}
