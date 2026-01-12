package main

import "fmt"

func main() {
	n := 5
	fmt.Println("A. Hasil dari Segitiga Bintang Terbalik:")
	segitigaBintangTerbalik(n)

	fmt.Println()

	fmt.Println("B. Hasil dari Segitiga Bintang Kanan:")
	segitigaBintangKanan(n)
}

func segitigaBintangTerbalik(n int) {

	for i := n; i >= 1; i-- {
		// cetak spasi
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		// cetak bintang
		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func segitigaBintangKanan(n int) {

	for i := 1; i <= n; i++ {
		// cetak spasi
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		// cetak bintang
		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}