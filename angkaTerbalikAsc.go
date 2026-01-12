package main

import "fmt"

func main() {
	n := 8
	segitigaBintangTerbalik(n)
}

func segitigaBintangTerbalik(n int) {
	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		// cetak angka ascending (kanan), mulai dari 2
		for j := 2; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
