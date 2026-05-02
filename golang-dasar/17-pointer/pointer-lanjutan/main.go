package main

import "fmt"

func heal(hp *int) {
	// tambah HP langsung di memori
	*hp += 20	// *hp = *hp + 20
	fmt.Println("Pemain disembuhkan +20!")
}

func attack(hp *int, damage int) {
	// kurangi HP langsung di memori
	*hp -= damage	// *hp = *hp - damage
	fmt.Printf("Pemain diserang! HP berkurang %d!\n", damage)
	if *hp <= 0 {
		fmt.Println("Game Over! Pemain kalah 😵")
	}
}

func main() {
	hp := 50 // HP awal
	fmt.Println("HP Awal:", hp)

	heal(&hp)	// kirim alamat hp
	fmt.Println("HP Sekarang:", hp)

	attack(&hp, 30)	// kirim alamat hp dan kasih damage = 30
	fmt.Println("HP Sekarang:", hp)

	attack(&hp, 50)	// kirim alamat hp dan kasih damage = 50
	fmt.Println("HP Sekarang:", hp)
}