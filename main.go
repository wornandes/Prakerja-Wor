package main

import (
	"fmt"
)

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	if n%2 == 0 {
// 		return false
// 	}
// 	for i := 3; i*i <= n; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var num int
// 	fmt.Print("Masukkan sebuah angka: ")
// 	fmt.Scan(&num)

// 	if isPrime(num) {
// 		fmt.Printf("%d adalah bilangan prima\n", num)
// 	} else {
// 		fmt.Printf("%d bukan bilangan prima\n", num)
// 	}
// }

func main() {
	var num int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	if num%7 == 0 {
		fmt.Printf("%d adalah kelipatan 7\n", num)
	} else {
		fmt.Printf("%d bukan kelipatan 7\n", num)
	}
}

// func hitungLuasTrapesium(alas, atas, tinggi float64) float64 {
// 	luas := 0.5 * (alas + atas) * tinggi
// 	return luas
// }

// func main() {
// 	var alas, atas, tinggi float64

// 	fmt.Print("Masukkan panjang alas: ")
// 	fmt.Scan(&alas)

// 	fmt.Print("Masukkan panjang atas: ")
// 	fmt.Scan(&atas)

// 	fmt.Print("Masukkan tinggi: ")
// 	fmt.Scan(&tinggi)

// 	luas := hitungLuasTrapesium(alas, atas, tinggi)

// 	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
// }
