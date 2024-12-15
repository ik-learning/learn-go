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
		{"8 gets converted to Roman", 8, "VIII"},
		{"9 gets converted to Roman", 9, "IX"},
		{"10 gets converted to Roman", 10, "X"},
		{"14 gets converted to Roman", 14, "XIV"},
		{"18 gets converted to Roman", 18, "XVIII"},
		{"20 gets converted to Roman", 20, "XX"},
		{"39 gets converted to Roman", 39, "XXXIX"},
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

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, roman := range allRomanNumerals {
		for arabic >= roman.Value {
			result.WriteString(roman.Symbol)
			arabic -= roman.Value
		}
	}
	return result.String()
}
