package parser

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/s-nix/mk2i18n/message"
)

func ToJson(messages []message.Message) (string, error) {
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

func DecodeJsonFile(path string, v any) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	return json.NewDecoder(fp).Decode(v)
}

func FromJson(inputPath string) ([]message.Message, error) {
	var messages []message.Message
	var Data map[string]any
	err := DecodeJsonFile(inputPath, &Data)
	if err != nil {
		return nil, err
	}
	flattenDataToMessages(Data, &messages, "")
	if len(messages) == 0 {
		return nil, nil
	}
	return messages, nil
}
