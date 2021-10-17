package transform

import (
	"strings"
	"unicode"
)

func Reverse(word string) string {
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		words[i] = reverseEachWord(words[i])
		// print(words)
	}
	return strings.Join(words, " ")
}
func reverseEachWord(word string) string {
	rns := []rune(word)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		if unicode.IsUpper(rns[i]) {
			rns[i], rns[j] = unicode.ToUpper(rns[j]), unicode.ToLower(rns[i])
		} else {
			rns[i], rns[j] = rns[j], rns[i]
		}
	}
	return string(rns)
}
