package parser

import (
	"github.com/s-nix/mk2i18n/message"
)

func ToToml(messages []message.Message) (string, error) {
	var result = ""
	for _, msg := range messages {
		tomlBytes, err := msg.MarshalTOML()
		if err != nil {
			return "", err
		}
		result += string(tomlBytes) + "\n"
	}
	return result, nil
}
