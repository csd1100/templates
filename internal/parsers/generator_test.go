package parsers_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/csd1100/templates/internal/parsers"
)

func TestGenerator(t *testing.T) {
	expected_files := []string{"../../tests/data/expected1", "../../tests/data/expected2"}
	expected_data := []string{}
	for _, expected_file := range expected_files {
		data, err := os.ReadFile(expected_file)
		if err != nil {
			t.Errorf("Unable to read expected files %s due to error %s", expected_file, err.Error())
		}
		expected_data = append(expected_data, string(data))
	}

	cases := []struct {
		name                  string
		config                parsers.Config
		templateFiles         parsers.TemplateFiles
		actual_files_to_check []string
		actual_file_present   bool
		expected_error        error
	}{
		{
			name: "generates valid templates when valid input",
			config: parsers.Config{
				ConfigFile:      "template-files.json",
				SourceDirectory: "../../tests/data/",
				TargetDirectory: "../../tests/generated/",
			},
			templateFiles: parsers.TemplateFiles{
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
			actual_files_to_check: []string{"../../tests/generated/testsource1.tmpl", "../../tests/generated/testsource2.tmpl"},
			actual_file_present:   true,
			expected_error:        nil,
		},
		{
			name: "generates valid templates when nested input",
			config: parsers.Config{
				ConfigFile:      "../../tests/data/nested-valid-template-files.json",
				SourceDirectory: "../../tests/data/",
				TargetDirectory: "../../tests/generated/",
			},
			templateFiles: parsers.TemplateFiles{
				Files: []parsers.TemplateFile{
					{
						Real:  "./testsource1",
						Template: "./dir1/testsource1.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
							"___packageName___": "{{ .packageName }}",
						},
					},
					{
						Real:  "./testsource2",
						Template: "./dir2/testsource2.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
						},
					},
				},
			},
			actual_files_to_check: []string{"../../tests/generated/dir1/testsource1.tmpl", "../../tests/generated/dir2/testsource2.tmpl"},
			actual_file_present:   true,
			expected_error:        nil,
		},
		{
			name: "returns error when invalid first source",
			config: parsers.Config{
				ConfigFile:      "../../tests/data/first-invalid-template-files.json",
				SourceDirectory: "../../tests/data/",
				TargetDirectory: "../../tests/generated/",
			},
			templateFiles: parsers.TemplateFiles{
				Files: []parsers.TemplateFile{
					{
						Real:  "./invalid_testsource1",
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
			actual_files_to_check: []string{"../../tests/generated/testsource1.tmpl", "../../tests/generated/testsource2.tmpl"},
			actual_file_present:   false,
			expected_error:        fmt.Errorf("no such file or directory"),
		},
		{
			name: "returns error when invalid second source",
			config: parsers.Config{
				ConfigFile:      "../../tests/data/second-invalid-template-files.json",
				SourceDirectory: "../../tests/data/",
				TargetDirectory: "../../tests/generated/",
			},
			templateFiles: parsers.TemplateFiles{
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
						Real:  "./invalid_testsource2",
						Template: "./testsource2.tmpl",
						Replacements: map[string]string{
							"___projectName___": "{{ .projectName }}",
						},
					},
				},
			},
			actual_files_to_check: []string{"../../tests/generated/testsource1.tmpl", "../../tests/generated/testsource2.tmpl"},
			actual_file_present:   false,
			expected_error:        fmt.Errorf("no such file or directory"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			t.Cleanup(func() {
				err := os.RemoveAll("../../tests/generated/")
				if err != nil {
					t.Errorf("Error while cleaning up tests/generated directory")
				}
			})

			err := parsers.Generate(&tc.config, &tc.templateFiles)
			fmt.Println(tc.name)
			fmt.Println(err)

			if err != nil {
				if !strings.Contains(err.Error(), tc.expected_error.Error()) {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, "error to contain '"+tc.expected_error.Error()+"'", err)
				}
			}
			if tc.actual_file_present {
				for index, actual := range tc.actual_files_to_check {
					actual_data, err := os.ReadFile(actual)
					if err != nil {
						t.Errorf("Unable to read generated file %v with error %s", actual, err.Error())
					}
					if string(actual_data) != expected_data[index] {
						t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, expected_data[index], actual_data)
					}
				}
			} else {
				for _, actual := range tc.actual_files_to_check {
					if _, err := os.Stat(actual); err == nil {
						t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, actual+" to be not present", actual+" is present")
					}
				}
			}
		})
	}
}
