package main

import (
	"fmt"
	"os"
)

type Profile struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func printBiodata(request Profile) {
	fmt.Println("Nama: ", request.Nama)
	fmt.Println("Alamat: ", request.Alamat)
	fmt.Println("Pekerjaan: ", request.Pekerjaan)
	fmt.Println("Alasan: ", request.Alasan)
}

profile := Profile{
	Nama:      "Andre",
	Alamat:    "Kebagusan",
	Pekerjaan: "Tidak ada",
	Alasan:    "Tertarik menjadi dev",
}
printBiodata(profile)



func main() {
	args := os.Args[1:]
	data := []profile
		profile
	fmt.Println(args)


	
}
