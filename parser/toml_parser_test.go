package parser

import (
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

	tomlOutput, err := ToToml(testMessages)
	assert.NoError(t, err)

	assert.Equal(t, expectedToml, tomlOutput)
}
