package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/s-nix/mk2i18n/converter"
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
		inFile  string
		outFile string
	)
	flag.StringVar(&inFile, "i", "", "Input file path. Supported formats are .json, .toml, .yaml, .xml, .properties, .csv")
	flag.StringVar(&outFile, "p", "", "Output file path. Supported formats are .json, .toml, .yaml.")
	flag.Parse()
	outPath, outFileName := filepath.Split(outFile)

	if outPath == "" {
		path, err := os.Getwd()
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Failed to get current working directory: %v\n", err)
			if err != nil {
				os.Exit(1)
			}
			os.Exit(2)
		}
		outPath = path
	}
	outType := filepath.Ext(outFileName)

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
	outFileInfo, err := os.Stat(outFile)
	if err == nil && !outFileInfo.IsDir() {
		_, err := fmt.Fprintf(os.Stderr, "Output path is a file, not a directory: %s\n", outFile)
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
	// Ensure the output file format is supported
	supported = false
	for _, ext := range SupportedOutputFormats {
		if outType == ext {
			supported = true
			break
		}
	}
	if !supported {
		_, err := fmt.Fprintf(os.Stderr, "Output file format not supported: %s\n", outType)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(2)
	}

	convertedString, err = converter.Convert(inFile, outFile)
}
