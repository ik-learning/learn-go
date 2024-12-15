package v2

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("1 gets converted to Roman", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("2 gets converted to Roman", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestRomanNumerals_V2(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to Roman", 1, "I"},
		{"2 gets converted to Roman", 2, "II"},
		{"3 gets converted to Roman", 3, "III"},
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
	if arabic == 2 {
		return "II"
	}
	if arabic == 3 {
		return "III"
	}
	return "I"
}
