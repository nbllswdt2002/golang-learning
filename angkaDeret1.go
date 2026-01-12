package main

import "fmt"

func main() {
	tampilkanPola(9)
	fmt.Println()
	tampilkanPola1(5)
}

func tampilkanPola(jumlahBaris int) {
	for i := 0; i < jumlahBaris; i++ {
		if i%2 == 0 {
			// Baris dengan angka genap
			for j := 2; j <= 8; j += 2 {
				if j == 2 {
					fmt.Print("- ")
				}
				fmt.Printf("%d - ", j)
			}
			fmt.Println()
		} else {
			// Baris dengan angka ganjil
			for j := 1; j <= 9; j += 2 {
				fmt.Printf("%d - ", j)
			}
			fmt.Println()
		}
	}
}

func tampilkanPola1(jumlahBaris int) {
	for i := 0; i < jumlahBaris; i++ {
		if i%2 == 0 {
			// Baris dengan angka genap
			for j := 2; j <= 4; j += 2 {
				if j == 2 {
					fmt.Print("- ")
				}
				fmt.Printf("%d - ", j)
			}
			fmt.Println()
		} else {
			// Baris dengan angka ganjil
			for j := 1; j <= 5; j += 2 {
				fmt.Printf("%d - ", j)
			}
			fmt.Println()
		}
	}
}
