# Templates

This repository contains templates for
[csd1100/init](https://github.com/csd1100/init) repository.

## Overview

### Templates setup

- Each template will be stored in branch of this repository. e.g. for rust projects
  there is branch named [rust](https://github.com/csd1100/templates/tree/rust).
- Each template branch should contain basic project setup along with a unit test
  and executable code.
- Each branch will have `templates/` directory which will hold go text templates
  that `init` will use to create projects.
- Each branch will also have one `template-files.json` which will hold information
  about templates for that project.

### Template Generation

- Templates can be generated using cli utility stored in main branch.
- The `template-generator` can be build using command:

```bash
go build -o ./build/template-generator ./cmd/template-generator/
```

- The template generator supports following flags:

```stdout
Usage:
  -c string
    Path or name of the config file to use for generating templates (default "template-files.json")
  -config string
    Path or name of the config file to use for generating templates (default "template-files.json")
  -s string
    Path to directory where files should be read from
  -source string
    Path to directory where files should be read from
  -t string
    Path to directory where templates should be generated
  -target string
    Path to directory where templates should be generated
  -v
    Print debug output
  -verbose
    Print debug output
```

- This utility requires `-s` to be passed.
- `-s` is source directory from which files will be read.
- `-t` is target directory where generated template will be stored. If not passed
  value will be same as `-s`
- `template-generator` will read config file specified in `-c` flag or
  if not defined it will read config from `source_directory/template-files.json`
- The config file will be stored in root of all branches named `template-files.json`.
- A sample `template-files.json`:

```json
{
  "template-files": [
    {
      "real": "./testsource1",
      "template": "./testsource1.tmpl",
      "replacements": {
        "___projectName___": "{{ .projectName }}",
        "___packageName___": "{{ .packageName }}"
      }
    },
    {
      "real": "./testsource2",
      "template": "./testsource2.tmpl",
      "replacements": {
        "___projectName___": "{{ .projectName }}"
      }
    }
  ]
}
```

- The template-generator will read file from `real`, and replace all occurrences
  of `___projectName___` with `{{ .projectName }}` and store generated file in `template`.
- The `init` project will read `template` and generate actual files at location
  of `real` depending on user input.
