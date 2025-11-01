package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/s-nix/mk2i18n/message"
)

type XMLFile struct {
	Data map[string]any
}

func (c *XMLFile) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	c.Data = map[string]any{}

	key := ""
	val := ""
	usedKeys := map[string]bool{}
	var multiFields []string
	i := 0
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			i = 0
			if key == "" {
				key = tt.Name.Local
			} else {
				key = key + "." + tt.Name.Local
			}
			indexedKey := key

			if _, exists := usedKeys[indexedKey]; exists {
				if !slices.Contains(multiFields, key) {
					multiFields = append(multiFields, key)
				}
				for {
					i++
					indexedKey = fmt.Sprintf("%s.%d", indexedKey, i)
					if _, exists := usedKeys[indexedKey]; !exists {
						break
					}
				}
			} else {
				indexedKey = fmt.Sprintf("%s.%d", indexedKey, i)
			}

			usedKeys[key] = true
			key = indexedKey

		case xml.CharData:
			val = string(tt)
			if strings.TrimSpace(val) == "" {
				continue
			}
			c.Data[key] = val

		case xml.EndElement:
			rTrailingNumber, err := regexp.Compile("\\.\\d+$")
			if err != nil {
				return err
			}
			key = rTrailingNumber.ReplaceAllString(key, "")
			i = 0
			split := strings.Split(key, ".")
			suffix := split[len(split)-1]
			key = strings.TrimSuffix(key, suffix)
			key = strings.TrimSuffix(key, ".")
			if tt.Name == start.Name {
				var keyNames []string
				var finalData = map[string]any{}
				for k, _ := range c.Data {
					keyNames = append(keyNames, k)
				}
				sort.Slice(keyNames, func(i, j int) bool {
					return keyNames[i] < keyNames[j]
				})
				for _, name := range keyNames {
					k := name
					elements := strings.Split(k, ".")
					var keyBuilder []string
					for i2, element := range elements {
						if element == "0" && i2 != 0 {
							prefix := strings.Join(keyBuilder, ".")
							// Check if the previous part is a multi-field
							if slices.Contains(multiFields, prefix) {
								keyBuilder = append(keyBuilder, element)
							} else {
								for i3, field := range multiFields {
									if strings.HasPrefix(field, prefix) {
										replacePrefix := prefix + ".0"
										newField := strings.Replace(field, replacePrefix, prefix, 1)
										multiFields = append(multiFields, newField)
										// Remove the old field
										multiFields = append(multiFields, newField)
										multiFields = append(multiFields[:i3], multiFields[i3+1:]...)
									}
								}
							}
						} else {
							keyBuilder = append(keyBuilder, element)
						}
					}
					finalKey := strings.Join(keyBuilder, ".")
					finalData[finalKey] = c.Data[name]
				}
				c.Data = finalData
				return nil
			}
		}
	}
}

func FromXML(inputPath string) ([]message.Message, error) {
	var messages []message.Message
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}
	data := XMLFile{}
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	FlattenDataToMessages(data.Data, &messages, "")
	if len(messages) == 0 {
		return nil, nil
	}
	return messages, nil
}
