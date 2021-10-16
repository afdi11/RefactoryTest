package wordcek

func Palindrome(kata string) bool {
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-i-1] {
			return false
		}
		// fmt.Println(i)
	}
	return true
}

func PalindromePrint(isPalindrome bool) string {
	if isPalindrome {
		return " # --> palindrome"
	} else {
		return " # --> not palindrome"
	}
}
