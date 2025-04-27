package romantointeger

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "III",
			expected: 3,
		},
		{
			input:    "IV",
			expected: 4,
		},
		{
			input:    "LVIII",
			expected: 58,
		},
		{
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			value := romanToInt(tt.input)

			if value != tt.expected {
				t.Errorf("value = %d, expected %d", value, tt.expected)
			}
		})
	}
}
