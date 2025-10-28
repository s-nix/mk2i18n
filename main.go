package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var SupportedInputFormats = []string{
	".json",
	".toml",
	".yaml",
	".xml",
	".properties",
	".csv",
}

var SupportedOutputFormats = []string{
	".json",
	".toml",
	".yaml",
}

func main() {
	var (
		inFile    string
		outPrefix string
		outType   string
		outPath   string
		locale    string
	)
	path, err := os.Getwd()
	flag.StringVar(&inFile, "i", "", "Input file path. Supported formats are .json, .toml, .yaml, .xml, .properties, .csv")
	flag.StringVar(&outPrefix, "p", "", "Output file prefix.")
	flag.StringVar(&outType, "o", SupportedOutputFormats[1], "Output file path. Supported formats are .json, .toml, .yaml")
	flag.StringVar(&outPath, "o", path, "Output file directory. If specified, output files will be saved in this directory with the given prefix and appropriate extensions. If not specified, output files will be saved in the current working directory.")
	flag.StringVar(&locale, "l", "en", "Locale to use for internationalization. Use IETF language tag format, e.g., 'en', 'fr', 'es'.")
	flag.Parse()
	// Ensure the input path is provided
	if inFile == "" {
		_, err := fmt.Fprintf(os.Stderr, "Please provide input file.\n")
		if err != nil {
			os.Exit(1)
		}
		os.Exit(2)
	}

	// Ensure the input path exists
	inFileInfo, err := os.Stat(inFile)
	if os.IsNotExist(err) {
		_, err := fmt.Fprintf(os.Stderr, "Input file does not exist: %s\n", inFile)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(2)
	}

	// Ensure the input path is a file, not a directory
	if inFileInfo.IsDir() {
		_, err := fmt.Fprintf(os.Stderr, "Input path is a directory, not a file: %s\n", inFile)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(2)
	}

	// Ensure the input file is a supported format
	supported := false
	inExt := filepath.Ext(inFile)
	for _, ext := range SupportedInputFormats {
		if inExt == ext {
			supported = true
			break
		}
	}
	if !supported {
		_, err := fmt.Fprintf(os.Stderr, "Input file format not supported: %s\n", inExt)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(2)
	}

	// If the output path is specified, but it is a file, exit with an error
	outPathInfo, err := os.Stat(outPath)
	if err == nil && !outPathInfo.IsDir() {
		_, err := fmt.Fprintf(os.Stderr, "Output path is a file, not a directory: %s\n", outPath)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(2)
	}

	// Ensure the output path exists. If not, create it.
	_, err = os.Stat(outPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(outPath, os.ModePerm)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Failed to create output directory: %s\n", outPath)
			if err != nil {
				os.Exit(1)
			}
			os.Exit(2)
		}
	}

}
