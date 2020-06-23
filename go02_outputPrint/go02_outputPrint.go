package main
/*
Output program ini masih ada yang salah, dan perlu perbaikan.
Semua tips utk perbaikan ada di video https://youtu.be/7mid-pK5onU
*/

import "fmt"

func main() {
	fmt.Println("Nama:", "TechJawir1", "\nUmur:", 27)
	fmt.Println("--------------------------------")
	fmt.Print("Nama: ", "TechJawir2", "\nUmur: ", 28)
	fmt.Println("--------------------------------")
	fmt.Printf("Nama: %s Umur: %v\n", "TechJawir3", 29)
}
