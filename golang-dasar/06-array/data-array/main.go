package main

import "fmt"

func main() {
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println("Data array:", angka)											// [10 20 30]
	fmt.Println("Index ke-1 awalnya begini:", angka[1]) 	// 20

	angka[1] = 80
	fmt.Println("Index ke-1 setelah diubah:", angka[1]) 	// 80
	fmt.Println("Hasilnya:", angka)												// [10 80 30]
}