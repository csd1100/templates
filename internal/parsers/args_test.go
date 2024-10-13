package parsers_test

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/csd1100/templates/internal/parsers"
)

func TestParse(t *testing.T) {
	configFilePath, err := filepath.Abs("../../tests/data/template-files.json")
	if err != nil {
		t.Errorf("Error while generating absolute path for config")
	}
	sourceDirectoryPath, err := filepath.Abs("../../tests/data/")
	if err != nil {
		t.Errorf("Error while generating absolute path for source directory")
	}
	targetDirectoryPath, err := filepath.Abs("../../tests/generated/")
	if err != nil {
		t.Errorf("Error while generating absolute path for target directory")
	}
	invalidPath, err := filepath.Abs("invalid")
	if err != nil {
		t.Errorf("Error while generating absolute path for invalid path")
	}
	invalidConfigPath, err := filepath.Abs(path.Join("../../tests/data/", "invalid"))
	if err != nil {
		t.Errorf("Error while generating absolute path for invalid config path")
	}

	cases := []struct {
		name          string
		init          func()
		expectedValue *parsers.Config
		expectedError error
	}{
		{
			name:          "returns error if empty args",
			expectedValue: nil,
			expectedError: fmt.Errorf("the parameter `s|source` is required"),
		},
		{
			name: "returns error if source not included",
			init: func() {
			},
			expectedValue: nil,
			expectedError: fmt.Errorf("the parameter `s|source` is required"),
		},
		{
			name: "returns source if target not included",
			init: func() {
				parsers.FSet.Set("s", "../../tests/data/")
			},
			expectedValue: &parsers.Config{
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: sourceDirectoryPath,
			},
			expectedError: nil,
		},
		{
			name: "returns valid config",
			init: func() {
				parsers.FSet.Set("c", "template-files.json")
				parsers.FSet.Set("v", "true")
				parsers.FSet.Set("s", "../../tests/data/")
				parsers.FSet.Set("t", "../../tests/generated/")
			},
			expectedValue: &parsers.Config{
				Verbose:         true,
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: targetDirectoryPath,
			},
			expectedError: nil,
		},
		{
			name: "returns valid config without specifying config file",
			init: func() {
				parsers.FSet.Set("v", "false")
				parsers.FSet.Set("s", "../../tests/data/")
				parsers.FSet.Set("t", "../../tests/generated/")
			},
			expectedValue: &parsers.Config{
				Verbose:         false,
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: targetDirectoryPath,
			},
			expectedError: nil,
		},
		{
			name: "returns valid config with full args",
			init: func() {
				parsers.FSet.Set("config", "template-files.json")
				parsers.FSet.Set("source", "../../tests/data/")
				parsers.FSet.Set("target", "../../tests/generated/")
			},
			expectedValue: &parsers.Config{
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: targetDirectoryPath,
			},
			expectedError: nil,
		},
		{
			name: "returns error if invalid source directory",
			init: func() {
				parsers.FSet.Set("config", "template-files.json")
				parsers.FSet.Set("source", "invalid")
				parsers.FSet.Set("target", "../../tests/generated/")
			},
			expectedValue: nil,
			expectedError: fmt.Errorf("unable to read %v", invalidPath),
		},
		{
			name: "returns error if invalid config file",
			init: func() {
				parsers.FSet.Set("config", "invalid")
				parsers.FSet.Set("source", "../../tests/data/")
				parsers.FSet.Set("target", "../../tests/generated/")
			},
			expectedValue: nil,
			expectedError: fmt.Errorf("unable to read %v", invalidConfigPath),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			oldArgs := os.Args
			os.Args = oldArgs[:1]

			if tc.init != nil {
				tc.init()
			}

			t.Cleanup(func() {
				parsers.FSet.Visit(func(f *flag.Flag) {
					f.Value.Set(f.DefValue)
				})
			})

			actual, err := parsers.ParseArgs()

			if err != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, tc.expectedError, err)
				}
				if actual != nil {
					t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, tc.expectedValue, actual)
				}
			} else {
				if !reflect.DeepEqual(*actual, *tc.expectedValue) {
					t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, *tc.expectedValue, *actual)
				}
				if err != nil {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, tc.expectedError, err)
				}
			}
			os.Args = oldArgs
		})
	}
}
