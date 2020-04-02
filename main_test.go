package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnderscoreToCamel(t *testing.T) {
	testData := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "empty",
			input:    "",
			expected: "",
		},
		{
			desc:     "happy path 1",
			input:    "apple",
			expected: "Apple",
		},
		{
			desc:     "happy path 2",
			input:    "big_apple",
			expected: "BigApple",
		},
		{
			desc:     "happy path 3",
			input:    "big_apple_",
			expected: "BigApple",
		},
		{
			desc:     "not underscore, just skip",
			input:    "BigApple",
			expected: "BigApple",
		},
		{
			desc:     "mixed 1",
			input:    "API_V2",
			expected: "APIV2",
		},
		{
			desc:     "mixed 2",
			input:    "deviceID_callAPIv2",
			expected: "DeviceIDCallAPIv2",
		},
	}

	for _, td := range testData {
		t.Run(td.desc, func(t *testing.T) {
			assert.Equal(t, td.expected, underscoreToCamel(td.input))
		})
	}
}
