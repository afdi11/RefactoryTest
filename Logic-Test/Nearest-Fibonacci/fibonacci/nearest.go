package fibonacci

func Nearest(arr []int) int {
	var i int
	for i = 1; recursion(i) < sum(arr); i++ {
	}
	return recursion(i) - sum(arr)
}
func recursion(n int) int {
	if n <= 1 {
		return n
	}
	return recursion(n-1) + recursion(n-2)
}

func sum(arr []int) int {
	result := 0
	for _, value := range arr {
		result += value
	}
	return result
}
