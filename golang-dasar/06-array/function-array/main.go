package main

import "fmt"

func main() {
	number := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Data array :", number)											// [10 20 30 40 50]
	fmt.Println("Jumlah Elemen :", len(number))							// 5 -> jumlah elemen array, diperoleh dari fungsi len()
	fmt.Println("Index ke-1 awalnya begini :", number[1])		// 20

	number[1] = 80
	fmt.Println("Index ke-1 setelah diubah :", number[1])	// 80
	fmt.Println("\nHasilnya seperti ini :")
	// Iterasi seluruh elemen array dengan range
	for index, value := range number {
		fmt.Printf("index ke-%v = %v\n", index, value) 		// cetak data array
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	fmt.Println("\nData array-1 :", arr1)		// [1, 2, 3]
	fmt.Println("Data array-2 :", arr2)			// [4, 5, 6]

	fmt.Println("arr1 == arr2 maka hasilnya:", arr1==arr2) 	// false
	fmt.Println("arr1 != arr2 maka hasilnya:", arr1!=arr2)	// true
}