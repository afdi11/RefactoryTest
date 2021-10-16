package main

import "fmt"

func main() {
	fmt.Print("Masukkan 2 angka : ")
	var angka1 int
	var angka2 int
	fmt.Scan(&angka1)
	fmt.Scan(&angka2)
	for angka1 = angka1 + 4; angka1 <= angka2; angka1 += 4 {
		if (angka1 + 4) >= angka2 {
			print(angka1)
		} else {
			print(angka1, ", ")
		}
	}
}
