package main

import "fmt"

type Profile struct {
	Nama		string
	Umur		int
	Alamat		string
	Pekerjaan	string
	Agama		string
}

func main () {
	profile := Profile{
		Nama: 		"Andre",
		Umur: 		28,
		Alamat:		"Kebagusan",
		Pekerjaan:	"Tidak ada",
		Agama:		"Kristen",
	}
	fmt.Println(profile)
}

