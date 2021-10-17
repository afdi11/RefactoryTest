package main

import (
	"ReverseWords/transform"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// kata := "Afdi Fauzul Bahar"
	scanner := bufio.NewScanner(os.Stdin)
	var kata string
	// var kata string
	if scanner.Scan() {
		kata = scanner.Text()
		// fmt.Printf("Input was: %q\n", line)
	}
	reverseWord := transform.Reverse(kata)
	fmt.Print(kata, "\nInto\n", reverseWord)
}
