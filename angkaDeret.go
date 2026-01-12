package main

import "fmt"

// struct untuk deret
type Deret struct {
	kRow int
	lCol int
}

// method untuk mencetak pola
func (d Deret) CetakPola() {
	for i := 1; i <= d.kRow; i++ {
		for j := 1; j <= d.lCol; j++ {
			if j%2 == 1 {
				fmt.Print(i) // posisi ganjil: cetak nomor baris
			} else {
				fmt.Print(d.kRow - i + 1) // posisi genap: cetak kebalikan
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("A. Hasil dari Deret Baris 9 Kolom 10:")

	deret1 := Deret{kRow: 9, lCol: 10}
	deret1.CetakPola()

	fmt.Println()
	fmt.Println("B. Hasil dari Deret Baris 5 Kolom 5:")

	deret2 := Deret{kRow: 5, lCol: 5}
	deret2.CetakPola()
}
