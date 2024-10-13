package parsers

import (
	"log"
	"os"
	"path"
	"strings"
)

func Generate(config *Config, data *TemplateFiles) error {
	var dirsToCreate []string
	filesToWrite := map[string][]byte{}

	for _, templateFile := range data.Files {
		content, err := os.ReadFile(path.Join(config.SourceDirectory, templateFile.Real))
		if err != nil {
			return err
		}

		target := path.Join(config.TargetDirectory, templateFile.Template)
		if config.Verbose {
			log.Printf("target -> %s", target)
		}

		targetDir := path.Dir(target)
		if config.Verbose {
			log.Printf("target directory -> %s", targetDir)
		}
		if _, err = os.Stat(targetDir); err != nil {
			dirsToCreate = append(dirsToCreate, targetDir)
		}

		text := string(content)
		for from, to := range templateFile.Replacements {
			if config.Verbose {
				log.Printf("from %s -> to %s for %s", from, to, target)
			}
			if strings.Contains(text, from) {
				text = strings.ReplaceAll(text, from, to)
			}
		}
		if config.Verbose {
			log.Printf("replaced text for %s ->\n%s", target, text)
		}

		filesToWrite[target] = []byte(text)
	}

	for _, targetDir := range dirsToCreate {
		err := os.MkdirAll(targetDir, 0744)
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
