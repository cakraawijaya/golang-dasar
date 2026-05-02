package main

import "fmt"

// contoh fungsi cleanup sederhana
func cleanup() {
	fmt.Println("Cleanup: Menutup resource...")
}

func bacaConfig(namaFile string) {
	// defer selalu dipanggil meskipun ada panic
	defer cleanup()

	// recover diletakkan dalam defer untuk menangkap panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error:", r)
		}
	}()

	if namaFile == "" {
		panic("Nama file konfigurasi tidak boleh kosong")
	}

	fmt.Println("Membaca konfigurasi dari", namaFile)
	// anggap berhasil membaca
}

func main() {
	bacaConfig("")
	fmt.Println("Program tetap berjalan setelah panic")
}