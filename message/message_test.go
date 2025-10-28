package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage_MarshalJSON(t *testing.T) {
	msg := &Message{
		ID:          "greeting",
		Description: "A friendly greeting",
		Other:       "Hello, World!",
	}

	jsonData, err := msg.MarshalJSON()
	assert.NoError(t, err, "Expected no error during JSON marshaling")

	expectedJSON := `{"greeting":{"description":"A friendly greeting","other":"Hello, World!"}}`
	assert.JSONEq(t, expectedJSON, string(jsonData), "JSON output did not match expected")
}

func TestMessage_MarshalJSON_EmptyFields(t *testing.T) {
	msg := &Message{
		ID:    "farewell",
		Other: "Goodbye!",
	}

	jsonData, err := msg.MarshalJSON()
	assert.NoError(t, err, "Expected no error during JSON marshaling")

	expectedJSON := `{"farewell":{"description":"","other":"Goodbye!"}}`
	assert.JSONEq(t, expectedJSON, string(jsonData), "JSON output did not match expected for empty fields")
}

func TestMessage_MarshalJSON_SpecialCharacters(t *testing.T) {
	msg := &Message{
		ID:          "special_chars",
		Description: "Message with special characters: \n\t\" ' \\ /",
		Other:       "This is a test with emojis 😊🚀",
	}

	jsonData, err := msg.MarshalJSON()
	assert.NoError(t, err, "Expected no error during JSON marshaling")

	expectedJSON := `{"special_chars":{"description":"Message with special characters: \n\t\" ' \\ /","other":"This is a test with emojis 😊🚀"}}`
	assert.JSONEq(t, expectedJSON, string(jsonData), "JSON output did not match expected for special characters")
}

func TestMessage_MarshalJSON_EmptyMessage(t *testing.T) {
	msg := &Message{}

	jsonData, err := msg.MarshalJSON()
	assert.NoError(t, err, "Expected no error during JSON marshaling")

	expectedJSON := `{"":{"description":"","other":""}}`
	assert.JSONEq(t, expectedJSON, string(jsonData), "JSON output did not match expected for empty message")
}

func TestMessage_MarshalTOML(t *testing.T) {
	msg := &Message{
		ID:          "greeting",
		Description: "A friendly greeting",
		Other:       "Hello, World!",
	}

	tomlData, err := msg.MarshalTOML()
	assert.NoError(t, err, "Expected no error during TOML marshaling")

	expectedTOML :=
		`[greeting]
description = "A friendly greeting"
other = "Hello, World!"
`
	assert.Equal(t, expectedTOML, string(tomlData), "TOML output did not match expected")
}

func TestMessage_MarshalTOML_EmptyFields(t *testing.T) {
	msg := &Message{
		ID:    "farewell",
		Other: "Goodbye!",
	}

	tomlData, err := msg.MarshalTOML()
	assert.NoError(t, err, "Expected no error during TOML marshaling")

	expectedTOML :=
		`[farewell]
description = ""
other = "Goodbye!"
`
	assert.Equal(t, expectedTOML, string(tomlData), "TOML output did not match expected for empty fields")
}

func TestMessage_MarshalTOML_SpecialCharacters(t *testing.T) {
	msg := &Message{
		ID:          "special_chars",
		Description: "Message with special characters: \n\t\" ' \\ /",
		Other:       "This is a test with emojis 😊🚀",
	}

	tomlData, err := msg.MarshalTOML()
	assert.NoError(t, err, "Expected no error during TOML marshaling")

	expectedTOML :=
		`[special_chars]
description = "Message with special characters: \n\t\" ' \\ /"
other = "This is a test with emojis 😊🚀"
`
	assert.Equal(t, expectedTOML, string(tomlData), "TOML output did not match expected for special characters")
}

func TestMessage_MarshalTOML_EmptyMessage(t *testing.T) {
	msg := &Message{}

	tomlData, err := msg.MarshalTOML()
	assert.NoError(t, err, "Expected no error during TOML marshaling")

	expectedTOML :=
		`[""]
description = ""
other = ""
`
	assert.Equal(t, expectedTOML, string(tomlData), "TOML output did not match expected for empty message")
}

func TestMessage_MarshalYAML(t *testing.T) {
	msg := &Message{
		ID:          "greeting",
		Description: "A friendly greeting",
		Other:       "Hello, World!",
	}

	yamlData, err := msg.MarshalYAML()
	assert.NoError(t, err, "Expected no error during YAML marshaling")

	expectedYAML :=
		`greeting:
  description: A friendly greeting
  other: Hello, World!
`
	assert.Equal(t, expectedYAML, string(yamlData.([]byte)), "YAML output did not match expected")
}

func TestMessage_MarshalYAML_EmptyFields(t *testing.T) {
	msg := &Message{
		ID:    "farewell",
		Other: "Goodbye!",
	}

	yamlData, err := msg.MarshalYAML()
	assert.NoError(t, err, "Expected no error during YAML marshaling")

	expectedYAML :=
		`farewell:
  description: ""
  other: Goodbye!
`
	assert.Equal(t, expectedYAML, string(yamlData.([]byte)), "YAML output did not match expected for empty fields")
}

func TestMessage_MarshalYAML_SpecialCharacters(t *testing.T) {
	msg := &Message{
		ID:          "special_chars",
		Description: "Message with special characters: \n\t\" ' \\ /",
		Other:       "This is a test with emojis 😊🚀",
	}

	yamlData, err := msg.MarshalYAML()
	assert.NoError(t, err, "Expected no error during YAML marshaling")

	expectedYAML :=
		`special_chars:
  description: "Message with special characters: \n\t\" ' \\ /"
  other: "This is a test with emojis \U0001F60A\U0001F680"
`
	assert.Equal(t, expectedYAML, string(yamlData.([]byte)), "YAML output did not match expected for special characters")
}

func TestMessage_MarshalYAML_EmptyMessage(t *testing.T) {
	msg := &Message{}

	yamlData, err := msg.MarshalYAML()
	assert.NoError(t, err, "Expected no error during YAML marshaling")

	expectedYAML :=
		`"":
  description: ""
  other: ""
`
	assert.Equal(t, expectedYAML, string(yamlData.([]byte)), "YAML output did not match expected for empty message")
}
