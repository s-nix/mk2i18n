package parser

import (
	"os"
	"testing"

	"github.com/s-nix/mk2i18n/message"
	"github.com/stretchr/testify/assert"
)

func TestToJson(t *testing.T) {
	testMessages := []message.Message{
		{
			ID:          "greeting",
			Description: "A greeting message",
			Other:       "Hello",
		},
		{
			ID:          "farewell",
			Description: "A farewell message",
			Other:       "Goodbye",
		},
	}
	expectedJson := `{
  "farewell": {
	"description": "A farewell message",
	"other": "Goodbye"
  },
  "greeting": {
	"description": "A greeting message",
	"other": "Hello"
  }
}`

	jsonOutput, err := ToJSON(testMessages)
	assert.NoError(t, err)

	assert.JSONEq(t, expectedJson, jsonOutput)
}

func TestFromJson(t *testing.T) {
	jsonInput := `{
		"something": "value",
		"greeting": {
			"array": [1, 2, 3],
			"description": "A greeting message",
			"other": 2,
			"extra": {
				"info": 3.14
			}
		},
		"nested": [
			{"name": "nested1"},
			{"name": "nested2"}
		],
		"farewell": {
			"description": "A farewell message",
			"other": 1
		}
	}`
	file, err := os.CreateTemp("", "test-*.json")
	assert.NoError(t, err)

	_, err = file.WriteString(jsonInput)
	assert.NoError(t, err)

	messages, err := FromJSON(file.Name())
	assert.NoError(t, err)

	expectedMessages := []message.Message{
		{
			ID:          "farewell.description",
			Description: "",
			Other:       "A farewell message",
		},
		{
			ID:          "farewell.other",
			Description: "",
			Other:       "1",
		},
		{
			ID:          "greeting.array.0",
			Description: "",
			Other:       "1",
		},
		{
			ID:          "greeting.array.1",
			Description: "",
			Other:       "2",
		},
		{
			ID:          "greeting.array.2",
			Description: "",
			Other:       "3",
		},
		{
			ID:          "greeting.description",
			Description: "",
			Other:       "A greeting message",
		},
		{
			ID:          "greeting.extra.info",
			Description: "",
			Other:       "3.14",
		},
		{
			ID:          "greeting.other",
			Description: "",
			Other:       "2",
		},
		{
			ID:          "nested.0.name",
			Description: "",
			Other:       "nested1",
		},
		{
			ID:          "nested.1.name",
			Description: "",
			Other:       "nested2",
		},
		{
			ID:          "something",
			Description: "",
			Other:       "value",
		},
	}
	assert.Equal(t, expectedMessages, messages)
}
