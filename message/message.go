package message

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Message struct {
	ID          string
	Description string
	Other       string
}

func (m *Message) BuildMap() map[string]interface{} {
	propName := m.ID
	result := map[string]interface{}{}
	result[propName] = map[string]string{
		"description": m.Description,
		"other":       m.Other,
	}
	return result
}

func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.BuildMap())
}

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
