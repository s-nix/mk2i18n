package parser

import (
	"os"
	"testing"

	"github.com/s-nix/mk2i18n/message"
	"github.com/stretchr/testify/assert"
)

func TestToToml(t *testing.T) {
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
	expectedToml := `[greeting]
description = "A greeting message"
other = "Hello"

[farewell]
description = "A farewell message"
other = "Goodbye"

`

	tomlOutput, err := ToTOML(testMessages)
	assert.NoError(t, err)

	assert.Equal(t, expectedToml, tomlOutput)
}

func TestFromToml(t *testing.T) {
	tomlInput := `something = "value"

[greeting]
array = [1, 2, 3]
description = "A greeting message"
other = 2

[greeting.extra]
info = 3.14

[[nested]]
name = "nested1"

[[nested]]
name = "nested2"

[farewell]
description = "A farewell message"
other = 1
`
	file, err := os.CreateTemp("", "test-*.toml")
	assert.NoError(t, err)

	_, err = file.WriteString(tomlInput)
	assert.NoError(t, err)

	expectedMessages := []message.Message{
		{
			ID:    "something",
			Other: "value",
		},
		{
			ID:    "greeting.array.0",
			Other: "1",
		},
		{
			ID:    "greeting.array.1",
			Other: "2",
		},
		{
			ID:    "greeting.array.2",
			Other: "3",
		},
		{
			ID:    "greeting.description",
			Other: "A greeting message",
		},
		{
			ID:    "greeting.other",
			Other: "2",
		},
		{
			ID:    "greeting.extra.info",
			Other: "3.14",
		},
		{
			ID:    "nested.0.name",
			Other: "nested1",
		},
		{
			ID:    "nested.1.name",
			Other: "nested2",
		},
		{
			ID:    "farewell.other",
			Other: "1",
		},
		{
			ID:    "farewell.description",
			Other: "A farewell message",
		},
	}

	messages, err := FromTOML(file.Name())
	assert.NoError(t, err)

	assert.Equal(t, len(expectedMessages), len(messages))

	for _, expectedMsg := range expectedMessages {
		found := false
		for _, msg := range messages {
			if msg.ID == expectedMsg.ID && msg.Other == expectedMsg.Other {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected message not found: %+v", expectedMsg)
	}
}
