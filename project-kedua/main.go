package main

import "fmt"

func main() {
	// Deklarasi variabel dan tipe data.
	/*
		Metode deklarasi variabel yang tipe datanya
		juga dituliskan disebut sebagai manifest typing.
	*/
	// Cara pertama deklarasi variabel dan tipe data
	var firstName string = "Himawan"

	// Cara kedua deklarasi variabel dan tipe data
	var lastName string
	lastName = "Az~"

	fmt.Printf("Halo %s %s! \n", firstName, lastName)

	// Memanggil fungsi dengan cara lain
	fmt.Println("Halo", firstName, lastName+"!")
	// Tanda + digunakan untuk menggabungkan 2 string (string concatenation)

	// Deklarasi variabel tanpa tipe data
	/*
		Metode deklarasi variabel yang tipe data-nya diketahui secara otomatis
		dari data/nilai disebut konsep type inference.
	*/
	var firstName1 string = "Himawan" // Manifest typing
	lastName1 := "Az~"                // Type inference

	fmt.Printf("Halo %s %s!\n", firstName1, lastName1)
}
