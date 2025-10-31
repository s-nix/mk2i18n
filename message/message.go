package message

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// Message represents a localization message with an ID, Description, and Other string as specified in ICU MessageFormat.
type Message struct {
	// ID is the identifier for the message.
	// It is used as the key in the output formats.
	ID string

	// Description provides additional context about the message.
	Description string

	// Other contains the actual message string in ICU MessageFormat.
	Other string
}

// BuildMap builds a map representation of the Message.
func (m *Message) BuildMap() map[string]interface{} {
	propName := m.ID
	result := map[string]interface{}{}
	result[propName] = map[string]string{
		"description": m.Description,
		"other":       m.Other,
	}
	return result
}

// MarshalJSON marshals the Message into JSON format.
func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.BuildMap())
}

// MarshalTOML marshals the Message into TOML format.
func (m *Message) MarshalTOML() ([]byte, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	encoder.Indent = ""
	err := encoder.Encode(m.BuildMap())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalYAML marshals the Message into YAML format.
func (m *Message) MarshalYAML() (interface{}, error) {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err := encoder.Encode(m.BuildMap())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
