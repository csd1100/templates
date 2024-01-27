package parsers

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func Generate(config *Config, data *TemplateFiles) error {
	dirsToCreate := []string{}
	filesToWrite := map[string][]byte{}

	for _, templateFile := range data.Files {
		content, err := os.ReadFile(path.Join(config.SourceDirectory, templateFile.Src))
		if err != nil {
			return err
		}

		target := path.Join(config.TargetDirectory, templateFile.Dest)
		fmt.Println(target)

		targetDir := path.Dir(target)
		fmt.Println(targetDir)
		if _, err = os.Stat(targetDir); err != nil {
			dirsToCreate = append(dirsToCreate, targetDir)
		}

		text := string(content)
		for from, to := range templateFile.Replacements {
			fmt.Println(from)
			fmt.Println(to)
			if strings.Contains(text, from) {
				text = strings.ReplaceAll(text, from, to)
			}
		}
		fmt.Println(text)

		filesToWrite[target] = []byte(text)
	}

	for _, targetDir := range dirsToCreate {
		err := os.MkdirAll(targetDir, 0644)
		if err != nil {
			return err
		}
	}

	for target, text := range filesToWrite {
		err := os.WriteFile(target, text, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
