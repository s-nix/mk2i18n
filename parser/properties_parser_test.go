package parser

import (
	"os"
	"testing"

	"github.com/magiconair/properties"
	"github.com/stretchr/testify/assert"
)

func TestFromProperties(t *testing.T) {
	// Create a temporary properties file for testing
	propsContent := `
greeting=Hello
farewell=Goodbye
`
	tmpFile, err := properties.LoadString(propsContent)
	assert.NoError(t, err)

	// Write the content to a temporary file
	tmpFilePath := "test.properties"
	// Get a writer to write the properties content
	file, err := os.Create(tmpFilePath)
	assert.NoError(t, err)
	_, err = tmpFile.Write(file, properties.UTF8)
	assert.NoError(t, err)

	// Test the FromProperties function
	messages, err := FromProperties(tmpFilePath)
	assert.NoError(t, err)
	assert.Len(t, messages, 2)

	expectedMessages := map[string]string{
		"greeting": "Hello",
		"farewell": "Goodbye",
	}

	for _, msg := range messages {
		expectedValue, exists := expectedMessages[msg.ID]
		assert.True(t, exists)
		assert.Equal(t, expectedValue, msg.Other)
	}

	// Clean up the temporary file
	err = file.Close()
	assert.NoError(t, err)
	err = os.Remove(tmpFilePath)
	assert.NoError(t, err)
}
