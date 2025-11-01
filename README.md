# mk2i18n

Convert localization source files into go-i18n message files (JSON, TOML, YAML).

mk2i18n reads common markup formats like Java .properties, JSON, TOML, YAML, and XML, then emits go-i18n compatible files where each flattened key becomes a message with `description` and `other` fields.

- Inputs:
  - `.properties`
  - `.json`
  - `.toml`
  - `.yaml/.yml`
  - `.xml`
- Outputs:
  - `.json`
  - `.toml`
  - `.yaml`

## Why

- Centralize diverse source formats for translation into a single go-i18n schema
- Keep message IDs stable via deterministic key flattening
- Simple CLI and importable Go package

## Install

With Go installed:

```cmd
go install github.com/s-nix/mk2i18n@latest
```

Build from source in this repo:

```cmd
go build -o mk2i18n.exe .
```

You can then run `mk2i18n.exe` (Windows) or `./mk2i18n` (other platforms).

## Usage (CLI)

Flags:
- -i string  Input file path. Supported: .json, .toml, .yaml, .yml, .xml, .properties
- -p string  Output file path. Supported: .json, .toml, .yaml

Examples:

```cmd
mk2i18n.exe -i ./strings.properties -p ./out/messages.yaml
mk2i18n.exe -i ./en.json          -p ./dist/en.yaml
mk2i18n.exe -i ./messages.toml    -p ./messages.yaml
mk2i18n.exe -i ./strings.xml      -p ./strings.json
mk2i18n.exe -i ./bundle.yaml      -p ./bundle.toml
```

Behavior:
- The tool validates the input and output extensions and paths.
- If the output directory does not exist, it will be created.
- If the exact output file already exists, the CLI will not overwrite it and will exit with an error. Provide a new filename or remove the existing file.

Exit codes are non-zero on error.

## What the output looks like

The go-i18n schema this tool writes is a flat map from message IDs to an object with description and other fields. For example, given nested inputs like:

```json
{
  "greeting": {
    "description": "A greeting message",
    "other": "Hello"
  },
  "farewell": {
    "description": "A farewell message",
    "other": "Goodbye"
  }
}
```

The YAML output becomes:

```yaml
greeting.description:
  description: ""
  other: A greeting message

greeting.other:
  description: ""
  other: Hello

farewell.description:
  description: ""
  other: A farewell message

farewell.other:
  description: ""
  other: Goodbye
```

JSON output is a single object with the same keys, pretty-printed:

```json
{
  "greeting.description": {
    "description": "",
    "other": "A greeting message"
  },
  "greeting.other": {
    "description": "",
    "other": "Hello"
  },
  "farewell.description": {
    "description": "",
    "other": "A farewell message"
  },
  "farewell.other": {
    "description": "",
    "other": "Goodbye"
  }
}
```

TOML output creates one table per flattened key:

```toml
["greeting.description"]
description = ""
other = "A greeting message"

["greeting.other"]
description = ""
other = "Hello"

["farewell.description"]
description = ""
other = "A farewell message"

["farewell.other"]
description = ""
other = "Goodbye"
```

Note: Entries are sorted lexicographically by message ID, so the order may differ from the input but is stable.

## Input formats and how they are flattened

mk2i18n normalizes all inputs to a nested map structure and then flattens it using dot-separated keys. The rules are:

- Objects/maps become dot-joined keys: parent.child.grandchild
- Arrays/slices are indexed: key.0, key.1, ...
- Mixed maps from YAML (map[any]any) are converted to map[string]any when keys are strings
- Primitive array items (e.g., ["a", "b"]) become messages `key.0`, `key.1` with the printed value as `other`
- Non-string scalar values are stringified
- After flattening, each leaf value becomes a message with:
  - ID: the flattened key
  - Other: the leaf value string
  - Description: empty by default (unless the input itself already had flattened description/other pairs)

Specific sources:

- .properties: each property `a.b.c=Value` becomes a message with ID `a.b.c` and other `Value`.
- JSON/TOML/YAML: nested documents are flattened according to the rules above.
- XML: element names form the path; repeated sibling elements are indexed; text content becomes the value.

## Programmatic usage (Go)

Import and call the high-level converter:

```go
package main

import (
  "log"
  "github.com/s-nix/mk2i18n/converter"
)

func main() {
  if err := converter.Convert("./in.yaml", "./out/messages.toml"); err != nil {
    log.Fatal(err)
  }
}
```

Or use the parsers/formatters directly:

```go
package main

import (
  "fmt"
  "github.com/s-nix/mk2i18n/parser"
)

func main() {
  msgs, err := parser.FromJSON("./in.json")
  if err != nil { panic(err) }
  yaml, err := parser.ToYAML(msgs)
  if err != nil { panic(err) }
  fmt.Println(yaml)
}
```

Message type (for reference):
- ID string
- Description string
- Other string

Each message marshals to JSON/TOML/YAML as described above.

## Examples by format

```cmd
mk2i18n.exe -i ./examples/en.properties -p ./build/en.yaml
mk2i18n.exe -i ./examples/en.json       -p ./build/en.toml
mk2i18n.exe -i ./examples/en.toml       -p ./build/en.json
mk2i18n.exe -i ./examples/en.xml        -p ./build/en.yaml
mk2i18n.exe -i ./examples/en.yaml       -p ./build/en.yaml
```

## Development

- Run tests:

```cmd
go test ./...
```

- Key packages:
  - `converter`: high-level `Convert(in, out)` that routes to format-specific parsers/formatters based on file extensions
  - `parser`: `FromJSON`, `FromTOML`, `FromYAML`, `FromXML`, `FromProperties` and `ToJSON`, `ToTOML`, `ToYAML`
  - `parser/data_flatten.go`: shared flattening logic
  - `message`: `Message` type plus JSON/TOML/YAML marshalers

## Troubleshooting

- Input file does not exist: ensure `-i` points to a file, not a directory
- Unsupported extension: check the input/output extensions listed above (CLI expects `.yaml`, not `.yml`)
- Will not overwrite output: if a file already exists at `-p`, delete it or choose a different path
- YAML keys not strings: YAML maps with non-string keys are ignored for those entries during flattening
- Ordering differences: output is sorted by key; compare structurally (e.g., JSON compare) rather than by line order

## License

GPL-3.0. See [LICENSE](./LICENSE).

## Acknowledgements

- [go-yaml (v3)](https://pkg.go.dev/gopkg.in/yaml.v3) for YAML encoding/decoding
- [BurntSushi/toml](https://github.com/BurntSushi/toml) for TOML encoding/decoding
- [magiconair/properties](https://github.com/magiconair/properties) for .properties parsing
- [stretchr/testify](https://github.com/stretchr/testify) for the test suite

