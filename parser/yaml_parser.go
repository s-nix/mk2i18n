package parser

import (
	"os"

	"github.com/s-nix/mk2i18n/message"
	"gopkg.in/yaml.v3"
)

func ToYAML(messages []message.Message) (string, error) {
	var result = ""
	for _, msg := range messages {
		bytes, err := msg.MarshalYAML()
		if err != nil {
			return "", err
		}
		result += string(bytes.([]byte)) + "\n"
	}
	return result, nil
}

func FromYAML(inputPath string) ([]message.Message, error) {
	var messages []message.Message
	var data map[string]any
	err := DecodeYAMLFile(inputPath, &data)
	if err != nil {
		return nil, err
	}
	FlattenDataToMessages(data, &messages, "")
	if len(messages) == 0 {
		return nil, nil
	}
	return messages, nil
}

func DecodeYAMLFile(path string, v any) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	return yaml.NewDecoder(fp).Decode(v)
}
