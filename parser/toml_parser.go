package parser

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/s-nix/mk2i18n/message"
)

func ToToml(messages []message.Message) (string, error) {
	var result = ""
	for _, msg := range messages {
		bytes, err := msg.MarshalTOML()
		if err != nil {
			return "", err
		}
		result += string(bytes) + "\n"
	}
	return result, nil
}

func FromToml(inputPath string) ([]message.Message, error) {
	var messages []message.Message
	var Data map[string]any
	_, err := toml.DecodeFile(inputPath, &Data)
	if err != nil {
		return nil, err
	}
	flattenDataToMessages(Data, &messages, "")
	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages found in TOML file")
	}
	return messages, nil
}
