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
	for i := arabic; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}
	return result.String()
}
