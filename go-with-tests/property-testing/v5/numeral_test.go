package v3

import (
	"strings"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to Roman", 1, "I"},
		{"2 gets converted to Roman", 2, "II"},
		{"3 gets converted to Roman", 3, "III"},
		{"4 gets converted to Roman", 4, "IV"},
		{"5 gets converted to Roman", 5, "V"},
		{"6 gets converted to Roman", 6, "VI"},
		{"7 gets converted to Roman", 7, "VII"},
	}
	for _, tt := range cases {
		t.Run(tt.Description, func(t *testing.T) {
			got := ConvertToRoman(tt.Arabic)
			if got != tt.Want {
				t.Errorf("got %s, want %s", got, tt.Want)
			}
		})
	}
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for arabic > 0 {
		switch {
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}
	return result.String()
}
