package parser

import (
	"os"
	"testing"

	"github.com/s-nix/mk2i18n/message"
	"github.com/stretchr/testify/assert"
)

func TestToYAML(t *testing.T) {
	messages := []message.Message{
		{
			ID:          "greeting",
			Description: "A friendly greeting",
			Other:       "Hello, World!",
		},
		{
			ID:          "farewell",
			Description: "A friendly farewell",
			Other:       "Goodbye, World!",
		},
	}

	yamlStr, err := ToYAML(messages)
	assert.NoError(t, err, "Expected no error from ToYAML")

	expectedYAML := `greeting:
  description: A friendly greeting
  other: Hello, World!

farewell:
  description: A friendly farewell
  other: Goodbye, World!

`

	assert.Equal(t, expectedYAML, yamlStr, "YAML output did not match expected")
}

func TestFromYAML(t *testing.T) {
	yamlContent := `
greeting:
  description: A friendly greeting
  other: Hello, World!

farewell:
  description: A friendly farewell
  other: Goodbye, World!
`

	tmpFile := "test_messages.yaml"
	err := os.WriteFile(tmpFile, []byte(yamlContent), 0644)
	assert.NoError(t, err, "Expected no error writing temporary YAML file")
	defer os.Remove(tmpFile)
	messages, err := FromYAML(tmpFile)
	assert.NoError(t, err, "Expected no error from FromYAML")

	expectedMessages := []message.Message{
		{
			ID:          "farewell.description",
			Description: "",
			Other:       "A friendly farewell",
		},
		{
			ID:          "farewell.other",
			Description: "",
			Other:       "Goodbye, World!",
		},
		{
			ID:          "greeting.description",
			Description: "",
			Other:       "A friendly greeting",
		},
		{
			ID:          "greeting.other",
			Description: "",
			Other:       "Hello, World!",
		},
	}

	assert.Equal(t, expectedMessages, messages, "Parsed messages did not match expected")
}
