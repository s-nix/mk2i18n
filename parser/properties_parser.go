package parser

import (
	"github.com/magiconair/properties"
	"github.com/s-nix/mk2i18n/message"
)

func FromProperties(inputPath string) ([]message.Message, error) {
	props, err := properties.LoadFile(inputPath, properties.UTF8)
	if err != nil {
		return nil, err
	}

	var messages []message.Message
	for _, key := range props.Keys() {
		value, ok := props.Get(key)
		if !ok {
			continue
		}
		msg := message.Message{
			ID:    key,
			Other: value,
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
