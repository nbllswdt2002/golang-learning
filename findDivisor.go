package main

import (
	"fmt"
)

func main() {

	//1. Deklarasi Variabel menggunakan findDivisor
	fmt.Println("=== findDivisor ===")
	fmt.Println()

	//A. Bilangan 6
	num1 := 6
	divisors1 := findDivisor(num1)
	fmt.Printf("1. Hasil Pembagian dari %d: ", num1)
	for _, d := range divisors1 {
		fmt.Print(d, " ")
	}

	fmt.Println()
	fmt.Println()

	//B. Bilangan 24
	num2 := 24
	divisors2 := findDivisor(num2)
	fmt.Printf("2. Hasil Pembagian dari %d: ", num2)
	for _, d := range divisors2 {
		fmt.Print(d, " ")
	}
	fmt.Println()
	fmt.Println()

	//C. Bilangan 7
	num3 := 7
	divisors3 := findDivisor(num3)
	fmt.Printf("3. Hasil Pembagian dari %d: ", num3)
	for _, d := range divisors3 {
		fmt.Print(d, " ")
	}
	fmt.Println()
	fmt.Println()

}

// 1. Tampilkan sejumlah bilangan pembagi n number - findDivisor
func findDivisor(n int) []int {
	divisors := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
