package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Mengambil atau memisahkan angka-angka (digit) ===")
	fmt.Println()

	fmt.Print("1. Hasil extractDigit(12234):     ")
	extractDigit(12234)
	fmt.Println()

	fmt.Print("2. Hasil extractDigit(5432):     ")
	extractDigit(5432)
	fmt.Println()

	fmt.Print("3. Hasil extractDigit(1278):     ")
	extractDigit(1278)
	fmt.Println()
}

func extractDigit(n int) {
	// Konversi ke string
	str := fmt.Sprintf("%d", n)

	// Tampilkan digit dari belakang ke depan
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]), " ")
	}
	fmt.Println()
}
