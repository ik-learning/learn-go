package reverse

import (
	"strings"
)

// Reverse first convert the string to a mutable array of runes
// perfrorm the reverse operation and re-cast to a string
func reverse(s string) string {
	 chars := []rune(s)
	 for i, j := 0, len(chars)-1; i<len(chars)/2; i,j = i+1, j-1 {
		 chars[i], chars[j] = chars[j], chars[i]
	 }
	return string(chars)
}
// reverse_words Reversing a string by word is a similar process.
// we convert the string into an array of strings where each entry is a word. Next we apply the normal reverse loop to that array.
func reverse_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}