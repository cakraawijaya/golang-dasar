package main

import "fmt"

// Struct
type PersegiPanjang struct {
	Panjang float64
	Lebar 	float64
}

// Method: Menghitung luas
func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}
func LuasBiasa(panjang, lebar int) int {
	return panjang * lebar
}

func main() {
	// Membuat object-1
	pp1 := PersegiPanjang{Panjang: 10, Lebar: 5}
	// Tampilkan data
	fmt.Println("=== Persegi Panjang 1 ===")
	fmt.Println("Panjang:", pp1.Panjang)
	fmt.Println("Lebar:", pp1.Lebar)
	fmt.Println("Luas:", pp1.Luas())

	// Membuat object-2
	pp2 := PersegiPanjang{Panjang: 4, Lebar: 2}
	// Tampilkan data
	fmt.Println("\n=== Persegi Panjang 2 ===")
	fmt.Println("Panjang:", pp2.Panjang)
	fmt.Println("Lebar:", pp2.Lebar)
	fmt.Println("Luas:", pp2.Luas())

	// Tampilkan data dari fungsi umum (biasa) bukan fungsi struct
	fmt.Println("\n=== Persegi Panjang 3 ===")
	fmt.Println("Luas:", LuasBiasa(2, 3))
}