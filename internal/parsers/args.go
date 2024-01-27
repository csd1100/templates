package parsers

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var FSet = flag.FlagSet{}

var configFile string
var sourceDirectory string
var targetDirectory string

type Config struct {
	ConfigFile      string
	SourceDirectory string
	TargetDirectory string
}

func getConfigFilePath(configFile, sourceDirectory string) (string, error) {
	if filepath.IsAbs(configFile) {
		return getAbsPath(configFile, false)
	} else {
		return getAbsPath(path.Join(sourceDirectory, configFile), false)
	}
}

func getAbsPath(path string, chkDir bool) (string, error) {
	if absPath, err := filepath.Abs(path); err == nil {
		if stat, err := os.Stat(absPath); err == nil {
			if !chkDir {
				return absPath, nil
			} else {
				if stat.IsDir() {
					return absPath, nil
				} else {
					return "", fmt.Errorf("%v is not a directory", absPath)
				}
			}
		} else {
			return "", fmt.Errorf("Unable to read %v", absPath)
		}
	} else {
		return "", fmt.Errorf("Unable to get absolute path of %v", path)
	}
}

func ParseArgs() (*Config, error) {
	FSet.StringVar(&configFile, "c", "template-files.json", "Path or name of the file to use for generating templates")
	FSet.StringVar(&configFile, "config", "template-files.json", "Path or name of the file to use for generating templates")
	FSet.StringVar(&sourceDirectory, "s", "", "Path to directory where files should be read from")
	FSet.StringVar(&sourceDirectory, "source", "", "Path to directory where files should be read from")
	FSet.StringVar(&targetDirectory, "t", "", "Path to directory where templates should be generated")
	FSet.StringVar(&targetDirectory, "target", "", "Path to directory where templates should be generated")

	err := FSet.Parse(os.Args[1:])
	if err != nil {
		return nil, fmt.Errorf("Unable to parse arguments")
	}

	if configFile == "" {
		return nil, fmt.Errorf("The parameter `c|config` is required")
	}

	if sourceDirectory == "" {
		return nil, fmt.Errorf("The parameter `s|source` is required")
	}

	if targetDirectory == "" {
		return nil, fmt.Errorf("The parameter `t|target` is required")
	}

	config := Config{}

	configFile, err = getConfigFilePath(configFile, sourceDirectory)
	if err != nil {
		return nil, err
	}

	sourceDirectory, err = getAbsPath(sourceDirectory, true)
	if err != nil {
		return nil, err
	}

	targetDirectory, err = getAbsPath(targetDirectory, true)
	if err != nil {
		return nil, err
	}

	config.ConfigFile = configFile
	config.SourceDirectory = sourceDirectory
	config.TargetDirectory = targetDirectory

	return &config, nil
}
