package parser

import (
	"fmt"
	"sort"

	"github.com/s-nix/mk2i18n/message"
)

// FlattenDataToMessages flattens a nested map[string]any structure into a slice of message.Message.
// Each key in the nested structure is concatenated with its parent keys using dot notation.
// The resulting messages are appended to the provided messages slice.
// The messages are sorted by their ID before returning.
func FlattenDataToMessages(data map[string]any, messages *[]message.Message, parent string) {
	for key, value := range data {
		if parent != "" {
			key = parent + "." + key
		}
		switch v := value.(type) {
		case map[string]any:
			FlattenDataToMessages(v, messages, key)

		case map[any]any:
			convertedMap := make(map[string]any)
			for k, val := range v {
				strKey, ok := k.(string)
				if ok {
					convertedMap[strKey] = val
				}
			}
			FlattenDataToMessages(convertedMap, messages, key)

		case []map[string]any:
			for i, item := range v {
				newKey := fmt.Sprintf("%s.%d", key, i)
				FlattenDataToMessages(item, messages, newKey)
			}

		case []map[any]any:
			for i, item := range v {
				newKey := fmt.Sprintf("%s.%d", key, i)
				convertedMap := make(map[string]any)
				for k, val := range item {
					strKey, ok := k.(string)
					if ok {
						convertedMap[strKey] = val
					}
				}
				FlattenDataToMessages(convertedMap, messages, newKey)
			}

		case []any:
			mapSliceValue, ok := value.([]map[string]any)
			if ok {
				for i, item := range mapSliceValue {
					newKey := fmt.Sprintf("%s.%d", key, i)
					FlattenDataToMessages(item, messages, newKey)
				}
				continue
			}
			for i, item := range v {
				key := fmt.Sprintf("%s.%d", key, i)
				valueMap, ok := item.(map[string]any)
				if ok {
					FlattenDataToMessages(valueMap, messages, key)
					continue
				}
				msg := message.Message{ID: key, Other: fmt.Sprintf("%v", item)}
				*messages = append(*messages, msg)
			}

		default:
			msg := message.Message{ID: key, Other: fmt.Sprintf("%v", v)}
			*messages = append(*messages, msg)
		}
	}
	sort.Slice(*messages, func(i, j int) bool {
		return (*messages)[i].ID < (*messages)[j].ID
	})
}
