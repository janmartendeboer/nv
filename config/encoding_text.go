/*
Package config implements an Encoder and Decoder for a text based representation of Variable.

For the following input:
	Variables{
		&Variable{
			Key: "HOME",
			Value: "C:\Users\Gopher",
		},
		&Variable{
			Key: "USERNAME",
			Value: "Gopher",
		},
	}

The Encoder will output:
	HOME=C:\Users\Gopher
	USERNAME=Gopher
*/
package config

import (
	"fmt"
	"strings"
)

type textEncoder struct {
	Encoder
}

// Allows to Encode Variable structs into a byte sequence.
func (e textEncoder) Encode(variables ...*Variable) ([]byte, error) {
	result := make([]string, len(variables))

	for i, v := range variables {
		result[i] = fmt.Sprintf("%s=%s", v.Key, v.Value)
	}

	return []byte(strings.Join(result, "\n")), nil
}

type textDecoder struct {
	Decoder
}

// Allows to Decode a byte sequence into a list of Variables.
func (d textDecoder) Decode(payload []byte) (Variables, error) {
	variables := make([]*Variable, 0)

	for _, line := range strings.Split(string(payload), "\n") {
		if len(line) == 0 {
			continue
		}

		components := strings.SplitN(line, "=", 2)

		if len(components) != 2 {
			return nil, IllFormattedVariable
		}

		variables = append(variables, &Variable{
			Key:   components[0],
			Value: components[1],
		})
	}

	return variables, nil
}

func init() {
	RegisterEncoding(
		"text",
		struct {
			textEncoder
			textDecoder
		}{},
	)
}
