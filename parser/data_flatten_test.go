package parser

import (
	"testing"

	"github.com/s-nix/mk2i18n/message"
	"github.com/stretchr/testify/assert"
)

func TestFlattenDataToMessages(t *testing.T) {
	data := map[string]any{
		"simple_key": "simple_value",
		"nested": map[string]any{
			"inner_key":   "inner_value",
			"inner_array": []any{"value1", "value2"},
		},
		"array_of_maps": []map[string]any{
			{"map_key1": "map_value1"},
			{"map_key2": "map_value2"},
		},
	}

	var messages []message.Message
	flattenDataToMessages(data, &messages, "")

	expectedMessages := []message.Message{
		{ID: "array_of_maps.0.map_key1", Other: "map_value1"},
		{ID: "array_of_maps.1.map_key2", Other: "map_value2"},
		{ID: "nested.inner_array.0", Other: "value1"},
		{ID: "nested.inner_array.1", Other: "value2"},
		{ID: "nested.inner_key", Other: "inner_value"},
		{ID: "simple_key", Other: "simple_value"},
	}

	assert.Equal(t, expectedMessages, messages)
}
