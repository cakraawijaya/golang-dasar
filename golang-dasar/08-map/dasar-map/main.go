package main

import "fmt"

func main() {
	// Membuat map dengan literal
	mahasiswa := map[string]int {
		// Insert data
		"triady": 90,
		"joko": 80,
		"yusup": 85,
	}

	// Tampilkan nilai mahasiswa
	fmt.Println("Nilai UAS milik Triady:", mahasiswa["triady"])
	fmt.Println("Nilai UAS milik Joko:", mahasiswa["joko"])
	fmt.Println("Nilai UAS milik Yusup:", mahasiswa["yusup"])
}