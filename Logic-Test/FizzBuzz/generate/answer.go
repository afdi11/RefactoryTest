package generate

import (
	"fmt"
	"strconv"
	"strings"
)

func Answer(angka int) []string {
	var answer []string
	// fmt.Println(angka)
	for i := 1; i <= angka; i++ {
		// fmt.Println(i, " ", angka)
		if (i%3 == 0) && (i%5 == 0) {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}

	}
	return answer
}
func Cetak(arr []string) {
	output := "[\"" + strings.Join(arr, `","`) + `"]`
	fmt.Println(output)
}
