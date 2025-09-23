package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var jumlah int

	fmt.Scanln(&a, &b, &c, &d, &e)
	jumlah = a + b + c + d + e
	fmt.Println("Hasil dari penjumlahan adalah:", jumlah)

}