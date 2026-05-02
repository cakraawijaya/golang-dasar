package main

import "fmt"

func ubahNama(nama *string) {
	*nama = "Budi" // ubah langsung ke alamat memori aslinya
}

func main() {
	nama := "Andi"
	ubahNama(&nama)
	fmt.Println(nama) // Output: Budi (berubah!)
}