package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Kombinasi FOR Loop di Go ===")

	// 1. For standar: mencetak angka 1 sampai 5
	fmt.Println("\n1. For standar (i = 1; i <= 5; i++):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 2. While-style: mencetak angka genap sampai 10
	fmt.Println("\n2. While-style (hanya kondisi):")
	j := 2
	for j <= 10 {
		fmt.Printf("%d ", j)
		j += 2	// j = j + 2
	}
	fmt.Println()

	// 3. Infinite loop dengan break: menghitung mundur dari 5
	fmt.Println("\n3. Infinite loop dengan break:")
	k := 5
	for {
		fmt.Printf("%d ", k)
		k--	// k = k - 1
		if k == 0 {
			break // keluar dari loop
		}
	}
	fmt.Println()

	// 4. Range loop: iterasi elemen slice
	fmt.Println("\n4. Range loop untuk slice:")
	buah := []string{"apel", "jeruk", "pisang"}
	for index, value := range buah {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}