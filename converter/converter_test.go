package converter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPropertiesToToml(t *testing.T) {
	// Prepare test data
	propertiesContent := `greeting=Hello
farewell=Goodbye
`
	expectedToml := `[greeting]
description = ""
other = "Hello"

[farewell]
description = ""
other = "Goodbye"

`

	// Write properties content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_properties_*.properties")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(propertiesContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.toml")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output TOML file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output TOML file")

	// Compare the output with expected TOML content
	assert.Equal(t, expectedToml, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}
