package main

import (
	"fmt"
	"palindrome/wordcek"
)

func main() {
	var Kata string = "kasur ini rusak"
	isPalindrome := wordcek.Palindrome(Kata)
	fmt.Println(Kata, wordcek.PalindromePrint(isPalindrome))
}
