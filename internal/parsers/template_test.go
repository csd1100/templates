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
		name           string
		templateFile   string
		expected_value *parsers.TemplateFiles
		expected_error error
	}{
		{
			name:           "returns error if file does not exist",
			templateFile:   "invalid",
			expected_value: nil,
			expected_error: fmt.Errorf("Unable to read config file '%v'", "invalid"),
		},
		{
			name:           "returns error if invalid json",
			templateFile:   "../../tests/data/invalid-json.json",
			expected_value: nil,
			expected_error: fmt.Errorf("Unable to parse config '%v', due to error:", "../../tests/data/invalid-json.json"),
		},
		{
			name:         "returns valid TemplateFiles",
			templateFile: "../../tests/data/template-files.json",
			expected_value: &parsers.TemplateFiles{
				Files: []parsers.TemplateFile{
					{
						Real:  "./testsource1",
						Template: "./testsource1.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
							"___packageName___": "{{ .packageName }}",
						},
					},
					{
						Real:  "./testsource2",
						Template: "./testsource2.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
						},
					},
				},
			},
			expected_error: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			actual, err := parsers.ParseTemplateFiles(tc.templateFile)

			if err != nil {
				if !strings.Contains(err.Error(), tc.expected_error.Error()) {
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
		})
	}
}
