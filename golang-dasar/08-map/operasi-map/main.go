package main

import "fmt"

func main() {
	// Membuat map dengan make
	mahasiswa := make(map[string]string)

	// Insert data
	mahasiswa["nama"] =  "triady"
	mahasiswa["kelas"] = "A"

	// Tampilkan nama dan kelas
	fmt.Println("Nama:", mahasiswa["nama"])
	fmt.Println("kelas:", mahasiswa["kelas"])

	// Update key
	mahasiswa["jurusan"] = "Informatika"
	fmt.Println("Jurusan:", mahasiswa["jurusan"])	// Tampilkan jurusan

	// Ambil data lalu Hapus
	jurusan := mahasiswa["jurusan"]
	delete(mahasiswa, "jurusan")
	
	// Cek apakah data masih ada
	_, exists := mahasiswa["jurusan"]

	// Tampilkan proses penghapusan
	fmt.Println("\n>> Jurusan:", jurusan, "akan dihapus dari data induk...")
	fmt.Println(">> Data tersebut berhasil dihapus !")
	fmt.Println(">> Jurusan:", jurusan, "sudah tidak ada ?", !exists)

	// Tampilkan ulang datanya
	fmt.Println("\nMenampilkan ulang data:")
	fmt.Println("Nama:", mahasiswa["nama"])
	fmt.Println("kelas:", mahasiswa["kelas"])
}