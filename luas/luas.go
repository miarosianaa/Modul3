package main

import "fmt"

//Fungsi untuk menghitung luas persegi panjang
func hitungLuas(panjang float64, lebar float64) float64 {
	return panjang * lebar
}

func main() {
	//Mendefinisikan panjang dan lebar
	var panjang, lebar float64

	fmt.Print("Masukan panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukan lebar : ")
	fmt.Scan(&lebar)

	//Memanggil fungsi hitungLuas dan menampilkan hasilnya
	luas := hitungLuas(panjang, lebar)
	fmt.Printf("Luas persegi panjang: %2f\n", luas)
}
