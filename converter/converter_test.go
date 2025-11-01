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
	yamlContent = `greeting:
  description: A greeting message
  other: Hello

farewell:
  description: A farewell message
  other: Goodbye
`

	expectedJSON = `{
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
	expectedTOML = `["greeting.description"]
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
	expectedAltTOML = `["farewell.description"]
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
	expectedYAML = `greeting.description:
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

`
	expectedAltYAML = `farewell.description:
  description: ""
  other: A farewell message

farewell.other:
  description: ""
  other: Goodbye

greeting.description:
  description: ""
  other: A greeting message

greeting.other:
  description: ""
  other: Hello

`
)

// To JSON tests
func TestConvertPropertiesToJSON(t *testing.T) {
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
	assert.JSONEq(t, expectedJSON, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertTOMLToJSON(t *testing.T) {
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
	assert.JSONEq(t, expectedJSON, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertJSONToJSON(t *testing.T) {
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
	assert.JSONEq(t, expectedJSON, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertXMLToJSON(t *testing.T) {
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
	assert.JSONEq(t, expectedJSON, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertYAMLToJSON(t *testing.T) {
	// Write YAML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.yaml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
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
	assert.JSONEq(t, expectedJSON, string(outputData), "JSON output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

// To TOML tests
func TestConvertPropertiesToTOML(t *testing.T) {
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
	assert.Equal(t, expectedTOML, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertTOMLToTOML(t *testing.T) {
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
	assert.Equal(t, expectedAltTOML, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)
	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertJSONToTOML(t *testing.T) {
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
	assert.Equal(t, expectedAltTOML, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertXMLToTOML(t *testing.T) {
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
	assert.Equal(t, expectedAltTOML, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertYAMLToTOML(t *testing.T) {
	// Write YAML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.yaml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
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
	assert.Equal(t, expectedAltTOML, string(outputData), "TOML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

// To YAML tests
func TestConvertPropertiesToYAML(t *testing.T) {
	// Write properties content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_properties_*.properties")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(propertiesContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.yaml")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output YAML file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output YAML file")

	// Compare the output with expected YAML content
	assert.Equal(t, expectedYAML, string(outputData), "YAML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertTOMLToYAML(t *testing.T) {
	// Write TOML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.toml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(tomlContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.yaml")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output YAML file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output YAML file")

	// Compare the output with expected YAML content
	assert.Equal(t, expectedAltYAML, string(outputData), "YAML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertJSONToYAML(t *testing.T) {
	// Write JSON content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.json")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(jsonContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.yaml")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output YAML file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output YAML file")

	// Compare the output with expected YAML content
	assert.Equal(t, expectedAltYAML, string(outputData), "YAML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertXMLToYAML(t *testing.T) {
	// Write XML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.xml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(xmlContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.yaml")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output YAML file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output YAML file")

	// Compare the output with expected YAML content
	assert.Equal(t, expectedAltYAML, string(outputData), "YAML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}

func TestConvertYAMLToYAML(t *testing.T) {
	// Write YAML content to a temporary file
	tmpFile, err := os.CreateTemp("", "test_input_*.yaml")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove input temporary file")
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(yamlContent)
	assert.NoError(t, err)

	tmpOutputFile, err := os.CreateTemp("", "test_output_*.yaml")
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err, "Failed to remove output temporary file")
	}(tmpOutputFile.Name())

	err = Convert(tmpFile.Name(), tmpOutputFile.Name())
	assert.NoError(t, err, "Conversion failed")

	// Read the output YAML file
	outputData, err := os.ReadFile(tmpOutputFile.Name())
	assert.NoError(t, err, "Failed to read output YAML file")

	// Compare the output with expected YAML content
	assert.Equal(t, expectedAltYAML, string(outputData), "YAML output did not match expected")
	err = tmpFile.Close()
	assert.NoError(t, err)

	err = tmpOutputFile.Close()
	assert.NoError(t, err)
}
