package main

import "fmt"



func main() {
	var a, id int
	var name string
	var rating float64
	sales := [][]interface{}{{1,"PUBG", 4.0},{2,"FF", 5.0}}

		

    fmt.Printf("Pilih:\n")
	fmt.Printf("1. List Game:\n")
	fmt.Printf("2. Tambah Game :\n")
	fmt.Printf("3. Hapus Game :\n")
	fmt.Printf("4. Cari Game Berdasarkan Nama :\n")
	fmt.Printf("5. Game 3 terfavorit :\n")
	fmt.Printf("6. Game dengan rating diatas 4.0 :\n")
	fmt.Print("Pilih menu : ")
	fmt.Scan(&a)

	switch a{
		case 1:
			fmt.Println(sales)
		case 2:
			fmt.Print("ID : ")
			fmt.Scan(&id)
			fmt.Print("Nama Game : ")
			fmt.Scan(&name)
			fmt.Print("Rating : ")
			fmt.Scan(&rating)
			sales = append(sales, []interface{}{id,name, rating})
			fmt.Println(sales)
		case 3:
			fmt.Print("Belum Tersedia")
		case 4:
			fmt.Print("Belum Tersedia")
		case 5:
			fmt.Print("Belum Tersedia")
		case 6:
			fmt.Print("Belum Tersedia")
			
		default:fmt.Println("Menu Tidak Ada")
	}
}