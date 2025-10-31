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
//		.json       (JSON files)
//		.xml        (XML files)
//		.toml       (TOML files)
//
//		Output
//	 -------------
//		.json       (JSON file in go-i18n format)
//		.toml       (TOML file in go-i18n format)
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
	case ".json":
		messages, err = parser.FromJSON(inFile)
		if err != nil {
			return err
		}
	case ".xml":
		messages, err = parser.FromXML(inFile)
		if err != nil {
			return err
		}
	case ".toml":
		messages, err = parser.FromTOML(inFile)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported input file extension: %s", inExtension)
	}

	// Generate output file
	var output string
	switch outExtension {
	case ".json":
		output, err = parser.ToJSON(messages)
		if err != nil {
			return err
		}
	case ".toml":
		output, err = parser.ToTOML(messages)
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
