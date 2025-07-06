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
	lastName1 := "Az~"                // Type inference,
	// hanya bisa dilakukan didalam blok fungsi, misalnya main()

	fmt.Printf("Halo %s %s!\n", firstName1, lastName1)

	// Deklarasi multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	// Contoh lain multi variabel
	var four, fifth, sixth string = "empat", "lima", "enam"

	// Multi variabel lebih ringkas
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// Teknik type inference untuk multi variabel beda tipe data
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "Hello"

	fmt.Println(first, second, third)
	fmt.Println(four, fifth, sixth)
	fmt.Println(seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)

	/*
		Semua variabel yang dideklarasikan harus digunakan.
		Jika tidak, maka akan terjadi error saat di run.
		Solusi dari hal tersebut adalah dengan variabel underscore _
		Variabel underscore adalah reserved variable yang bisa digunakan
		untuk menampung nilai yang tidak dipakai. Bisa juga dikatakan sebagai
		keranjang sampah.
	*/
	// Contoh penggunaan reserved variable
	_ = "belajar Golang"
	_ = "Bismillah Golang"
	name, _ := "John", "Doe"

	println(name)

	// Deklarasi variabel menggunakan keyword new
	name1 := new(string)
	fmt.Println(name1)  //tipe pointer string, hanya akan muncul alamat memori (hexadecimal)
	fmt.Println(*name1) //untuk menampilkan nilai asli dengan asterisk sebagai dereference
}
