package parser

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/s-nix/mk2i18n/message"
)

// ToJSON converts a slice of message.Message objects into a pretty-printed JSON string.
// Each message.Message is marshaled to JSON and combined into a single JSON object.
func ToJSON(messages []message.Message) (string, error) {
	var result = ""
	for _, msg := range messages {
		bytes, err := msg.MarshalJSON()
		if err != nil {
			return "", err
		}
		value := string(bytes)
		result += strings.TrimPrefix(strings.TrimSuffix(value, "}"), "{") + ","
	}
	returnString := "{" + strings.TrimRight(result, ",") + "}"
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, []byte(returnString), "", "  ")
	if err != nil {
		return "", err
	}
	return prettyJson.String(), nil
}

func DecodeJSONFile(path string, v any) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	return json.NewDecoder(fp).Decode(v)
}

func FromJSON(inputPath string) ([]message.Message, error) {
	var messages []message.Message
	var data map[string]any
	err := DecodeJSONFile(inputPath, &data)
	if err != nil {
		return nil, err
	}
	FlattenDataToMessages(data, &messages, "")
	if len(messages) == 0 {
		return nil, nil
	}
	return messages, nil
}
