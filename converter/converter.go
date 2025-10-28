package converter

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/s-nix/mk2i18n/message"
	"github.com/s-nix/mk2i18n/parser"
)

// Convert converts the input file to the output file format as required by go-i18n.
// It supports various input and output formats based on file extensions.
// Currently supported formats are:
//
//		Input
//	 -------------
//		.properties (Java .properties files)
//
//
//		Output
//	 -------------
//		.toml (TOML files used by go-i18n)
func Convert(inFile string, outFile string) error {
	inExtension := filepath.Ext(inFile)
	outExtension := filepath.Ext(outFile)

	// Parse input file
	var messages []message.Message
	var err error
	switch inExtension {
	case ".properties":
		messages, err = parser.FromProperties(inFile)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported input file extension: %s", inExtension)
	}

	// Generate output file
	var output string
	switch outExtension {
	case ".toml":
		output, err = parser.ToToml(messages)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported output file extension: %s", outExtension)
	}

	err = os.WriteFile(outFile, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
