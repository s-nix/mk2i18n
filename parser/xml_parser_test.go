package parser

import (
	"os"
	"testing"

	"github.com/s-nix/mk2i18n/message"
	"github.com/stretchr/testify/assert"
)

func TestFromXML(t *testing.T) {
	testXML := `<resources>
	<greeting>
		<subGreeting>Hello</subGreeting>
		<subGreeting>World</subGreeting>
	</greeting>
	<greeting>
		<subGreeting>Bonjour</subGreeting>
		<subGreeting>Again</subGreeting>
	</greeting>
	<farewell>
		<subGreeting>Goodbye</subGreeting>
		<subGreeting>Moon</subGreeting>
	</farewell>
	<goodbye>Farewell</goodbye>
</resources>`

	tmpFile, err := os.CreateTemp("", "test_*.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(testXML)
	assert.NoError(t, err)
	tmpFile.Close()

	expectedMessages := []message.Message{
		{
			ID:          "farewell.subGreeting.0",
			Description: "",
			Other:       "Goodbye",
		},
		{
			ID:          "farewell.subGreeting.1",
			Description: "",
			Other:       "Moon",
		},
		{
			ID:          "goodbye",
			Description: "",
			Other:       "Farewell",
		},
		{
			ID:          "greeting.0.subGreeting.0",
			Description: "",
			Other:       "Hello",
		},
		{
			ID:          "greeting.0.subGreeting.1",
			Description: "",
			Other:       "World",
		},
		{
			ID:          "greeting.1.subGreeting.0",
			Description: "",
			Other:       "Bonjour",
		},
		{
			ID:          "greeting.1.subGreeting.1",
			Description: "",
			Other:       "Again",
		},
	}

	messages, err := FromXML(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
}
