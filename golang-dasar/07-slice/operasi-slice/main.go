package main

import "fmt"

func main() {
	// Membuat slice baru
	// Mengatur tipe data, panjang elemen, dan kapasitas
	slice1 := make([]int, 3, 5)		// [0 0 0 _ _] -> ada 2 ruang cadangan
	fmt.Println(slice1)						// [0 0 0] -> ada 3 ruang aktif
	fmt.Println("Panjang elemen :", len(slice1))
	fmt.Println("Kapasitas :", cap(slice1))

	// Menambahkan nilai pada bagian yang kosong
	slice1 = append(slice1, 10, 20)	// [0 0 0 10 20]
	fmt.Println("Setelah ditambahkan :", slice1)

	// Membuat slice baru dengan panjang elemen: 3
	slice2 := make([]int, 3)	// [0 0 0]

	// Menyalin isi slice1 ke slice2 (sebanyak yang muat di slice2)
	copied := copy(slice2, slice1)

	fmt.Println("Hasil salinan data :", slice2)
	fmt.Println("Jumlah elemen yang tersalin :", copied)

	// Ubah isi hasil copy
	slice2[0] = 2
	slice2[1] = 5
	slice2[2] = 8
	fmt.Println("Data diubah menjadi :", slice2)	// [2 5 8]

	// Ambil sisa dari slice1
	sisa := slice1[3:]

	// Hasil akhir setelah semua digabung
	result := append(slice2, sisa...)
	fmt.Println("Hasil akhir :", result)
}