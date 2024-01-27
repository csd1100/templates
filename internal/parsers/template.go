package parsers

import (
	"encoding/json"
	"fmt"
	"os"
)

type TemplateFile struct {
	Src          string            `json:"src"`
	Dest         string            `json:"dest"`
	Replacements map[string]string `json:"replacements"`
}

type TemplateFiles struct {
	Files []TemplateFile `json:"template-files"`
}

func ParseTemplateFiles(templateFile string) (*TemplateFiles, error) {
	var templateFiles TemplateFiles
	contents, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to read template files %v", templateFile)
	}

	err = json.Unmarshal(contents, &templateFiles)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse TemplateFiles %v due to error: %w", templateFile, err)
	}

	return &templateFiles, nil
}
