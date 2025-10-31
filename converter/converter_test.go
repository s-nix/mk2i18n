package converter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	propertiesContent = `greeting.description=A greeting message
greeting.other=Hello
farewell.description=A farewell message
farewell.other=Goodbye
`
	jsonContent = `{
  "greeting": {
	"description": "A greeting message",
	"other": "Hello"
  },
  "farewell": {
	"description": "A farewell message",
	"other": "Goodbye"
  }
}`
	tomlContent = `[greeting]
description = "A greeting message"
other = "Hello"

[farewell]
description = "A farewell message"
other = "Goodbye"
`
	xmlContent = `<resources>
	<greeting>
		<description>A greeting message</description>
		<other>Hello</other>
	</greeting>
	<farewell>
		<description>A farewell message</description>
		<other>Goodbye</other>
	</farewell>
</resources>`
	expectedJson = `{
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
}`
	expectedToml = `["greeting.description"]
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

`
	expectedAltToml = `["farewell.description"]
description = ""
other = "A farewell message"

["farewell.other"]
description = ""
other = "Goodbye"

["greeting.description"]
description = ""
other = "A greeting message"

["greeting.other"]
description = ""
other = "Hello"

`
)

func TestConvertPropertiesToJson(t *testing.T) {
	// Write properties content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_properties_*.properties")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(propertiesContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output JSON file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output JSON file")

	// Compare the output with expected JSON content
	assert.JSONEq(t, expectedJson, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertTomlToJson(t *testing.T) {
	// Write TOML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.toml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(tomlContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output JSON file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output JSON file")

	// Compare the output with expected JSON content
	assert.JSONEq(t, expectedJson, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertJsonToJson(t *testing.T) {
	// Write JSON content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.json")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(jsonContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output JSON file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output JSON file")

	// Compare the output with expected JSON content
	assert.JSONEq(t, expectedJson, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertXmlToJson(t *testing.T) {
	// Write XML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.xml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(xmlContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.json")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output JSON file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output JSON file")

	// Compare the output with expected JSON content
	assert.JSONEq(t, expectedJson, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertPropertiesToToml(t *testing.T) {
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

func TestConvertTomlToToml(t *testing.T) {
	// Write TOML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.toml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(tomlContent)
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
	assert.Equal(t, expectedAltToml, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)
	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertJsonToToml(t *testing.T) {
	// Write JSON content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.json")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(jsonContent)
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
	assert.Equal(t, expectedAltToml, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertXmlToToml(t *testing.T) {
	// Write XML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.xml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(xmlContent)
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
	assert.Equal(t, expectedAltToml, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}
