package parsers_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/csd1100/templates/internal/parsers"
)

func TestParseTemplateFile(t *testing.T) {
	cases := []struct {
		name          string
		templateFile  string
		expectedValue *parsers.TemplateFiles
		expectedError error
	}{
		{
			name:          "returns error if file does not exist",
			templateFile:  "invalid",
			expectedValue: nil,
			expectedError: fmt.Errorf("unable to read config file '%v'", "invalid"),
		},
		{
			name:          "returns error if invalid json",
			templateFile:  "../../tests/data/invalid-json.json",
			expectedValue: nil,
			expectedError: fmt.Errorf("unable to parse config '%v', due to error", "../../tests/data/invalid-json.json"),
		},
		{
			name:         "returns valid TemplateFiles",
			templateFile: "../../tests/data/template-files.json",
			expectedValue: &parsers.TemplateFiles{
				Files: []parsers.TemplateFile{
					{
						Real:     "./testsource1",
						Template: "./testsource1.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
							"___packageName___": "{{ .packageName }}",
						},
					},
					{
						Real:     "./testsource2",
						Template: "./testsource2.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
						},
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			actual, err := parsers.ParseTemplateFiles(tc.templateFile)

			if err != nil {
				if !strings.Contains(err.Error(), tc.expectedError.Error()) {
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
		})
	}
}
