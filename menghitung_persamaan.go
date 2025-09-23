package main

import "fmt"

func main() {
	var x int
	fmt.Print("masukkan nilai X:")
	fmt.Scanln(&x)
	var persamaan = 2/(float64(x)+5) + 5
	fmt.Println("hasil persamaan adalah:", persamaan)
}
