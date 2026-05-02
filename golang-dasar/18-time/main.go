package main

import (
	"fmt"
	"time"
)

// Fungsi untuk konversi waktu ke zona tertentu
func convertToZone(t time.Time, zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error memuat lokasi:", err)
		return t
	}
	return t.In(loc)
}

func main() {
	// 1. Ambil waktu saat ini dalam UTC
	nowUTC := time.Now().UTC()
	fmt.Println("Waktu UTC:", nowUTC)

	// 2. Konversi ke zona waktu lokal (Asia/Jakarta)
	jakartaTime := convertToZone(nowUTC, "Asia/Jakarta")
	fmt.Println("Waktu Jakarta:", jakartaTime)

	// 3. Konversi ke zona waktu lain yang menggunakan DST (misalnya New York)
	newYorkTime := convertToZone(nowUTC, "America/New_York")
	fmt.Println("Waktu New York:", newYorkTime)

	// 4. Menyimpan waktu (simpan ke penyimpanan ke DB)
	// Disarankan simpan waktu dalam UTC
	fmt.Println("Disimpan ke database (UTC):", nowUTC)
}