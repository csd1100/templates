package parsers

import (
	"errors"
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
var verbose bool

type Config struct {
	Help            bool
	Verbose         bool
	ConfigFile      string
	SourceDirectory string
	TargetDirectory string
}

func getAbsPath(path string) (string, error) {
	if absPath, err := filepath.Abs(path); err == nil {
		if _, err := os.Stat(absPath); err == nil {
			return absPath, nil
		} else {
			return "", fmt.Errorf("Unable to read %v", absPath)
		}
	} else {
		return "", fmt.Errorf("Unable to get absolute path of %v", path)
	}
}

func getConfigFilePath(configFile, sourceDirectory string) (string, error) {
	if path.Base(configFile) == configFile {
		return getAbsPath(path.Join(sourceDirectory, configFile))
	} else {
		return getAbsPath(configFile)
	}
}

func init() {
	FSet.StringVar(&configFile, "c", "template-files.json", "Path or name of the config file to use for generating templates")
	FSet.StringVar(&configFile, "config", "template-files.json", "Path or name of the config file to use for generating templates")
	FSet.StringVar(&sourceDirectory, "s", "", "Path to directory where files should be read from")
	FSet.StringVar(&sourceDirectory, "source", "", "Path to directory where files should be read from")
	FSet.StringVar(&targetDirectory, "t", "", "Path to directory where templates should be generated")
	FSet.StringVar(&targetDirectory, "target", "", "Path to directory where templates should be generated")
	FSet.BoolVar(&verbose, "v", false, "Print debug output")
	FSet.BoolVar(&verbose, "verbose", false, "Print debug output")
}

func ParseArgs() (*Config, error) {
	err := FSet.Parse(os.Args[1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return &Config{Help: true}, nil
		} else {
			return nil, fmt.Errorf("Unable to parse arguments")
		}
	}

	if configFile == "" {
		return nil, fmt.Errorf("The parameter `c|config` is required")
	}

	if sourceDirectory == "" {
		return nil, fmt.Errorf("The parameter `s|source` is required")
	}

	config := Config{}

	sourceDirectory, err = getAbsPath(sourceDirectory)
	if err != nil {
		return nil, err
	}

	configFile, err = getConfigFilePath(configFile, sourceDirectory)
	if err != nil {
		return nil, err
	}

	if targetDirectory == "" {
		targetDirectory = sourceDirectory
	} else {
		targetDirectory, err = filepath.Abs(targetDirectory)
		if err != nil {
			return nil, err
		}
	}

	config.Verbose = verbose
	config.ConfigFile = configFile
	config.SourceDirectory = sourceDirectory
	config.TargetDirectory = targetDirectory

	return &config, nil
}
