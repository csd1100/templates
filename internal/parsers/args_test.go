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
		name           string
		init           func()
		expected_value *parsers.Config
		expected_error error
	}{
		{
			name:           "returns error if empty args",
			expected_value: nil,
			expected_error: fmt.Errorf("The parameter `s|source` is required"),
		},
		{
			name: "returns error if source not included",
			init: func() {
			},
			expected_value: nil,
			expected_error: fmt.Errorf("The parameter `s|source` is required"),
		},
		{
			name: "returns source if target not included",
			init: func() {
				parsers.FSet.Set("s", "../../tests/data/")
			},
			expected_value: &parsers.Config{
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: sourceDirectoryPath,
			},
			expected_error: nil,
		},
		{
			name: "returns valid config",
			init: func() {
				parsers.FSet.Set("c", "template-files.json")
				parsers.FSet.Set("v", "true")
				parsers.FSet.Set("s", "../../tests/data/")
				parsers.FSet.Set("t", "../../tests/generated/")
			},
			expected_value: &parsers.Config{
				Verbose:         true,
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: targetDirectoryPath,
			},
			expected_error: nil,
		},
		{
			name: "returns valid config without specifying config file",
			init: func() {
				parsers.FSet.Set("v", "false")
				parsers.FSet.Set("s", "../../tests/data/")
				parsers.FSet.Set("t", "../../tests/generated/")
			},
			expected_value: &parsers.Config{
				Verbose:         false,
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: targetDirectoryPath,
			},
			expected_error: nil,
		},
		{
			name: "returns valid config with full args",
			init: func() {
				parsers.FSet.Set("config", "template-files.json")
				parsers.FSet.Set("source", "../../tests/data/")
				parsers.FSet.Set("target", "../../tests/generated/")
			},
			expected_value: &parsers.Config{
				ConfigFile:      configFilePath,
				SourceDirectory: sourceDirectoryPath,
				TargetDirectory: targetDirectoryPath,
			},
			expected_error: nil,
		},
		{
			name: "returns error if invalid source directory",
			init: func() {
				parsers.FSet.Set("config", "template-files.json")
				parsers.FSet.Set("source", "invalid")
				parsers.FSet.Set("target", "../../tests/generated/")
			},
			expected_value: nil,
			expected_error: fmt.Errorf("Unable to read %v", invalidPath),
		},
		{
			name: "returns error if invalid config file",
			init: func() {
				parsers.FSet.Set("config", "invalid")
				parsers.FSet.Set("source", "../../tests/data/")
				parsers.FSet.Set("target", "../../tests/generated/")
			},
			expected_value: nil,
			expected_error: fmt.Errorf("Unable to read %v", invalidConfigPath),
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
				if err.Error() != tc.expected_error.Error() {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, tc.expected_error, err)
				}
				if actual != nil {
					t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, tc.expected_value, actual)
				}
			} else {
				if !reflect.DeepEqual(*actual, *tc.expected_value) {
					t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, *tc.expected_value, *actual)
				}
				if err != nil {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, tc.expected_error, err)
				}
			}
			os.Args = oldArgs
		})
	}
}
