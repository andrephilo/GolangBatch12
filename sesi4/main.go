package main

import (
	"fmt"
	"os"
)

type Profile struct {
	Code      string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func printBiodata(request Profile) {
	fmt.Println("Code: ", request.Code)
	fmt.Println("Nama: ", request.Nama)
	fmt.Println("Alamat: ", request.Alamat)
	fmt.Println("Pekerjaan: ", request.Pekerjaan)
	fmt.Println("Alasan: ", request.Alasan)
}

var profile = Profile{
	Code:      "GLNG012ONL001",
	Nama:      "Rahmat Faisal",
	Alamat:    "dll",
	Pekerjaan: "dll",
	Alasan:    "dll",
}

func main() {
	args := os.Args[1:]
	data := []Profile{
		Profile{
			Code:      "GLNG012ONL001",
			Nama:      "Rahmat Faisal",
			Alamat:    "Depok",
			Pekerjaan: "dll",
			Alasan:    "dll",
		},
		Profile{
			Code:      "GLNG012ONL002",
			Nama:      "Irma Nur Afifah",
			Alamat:    "Kebayoran",
			Pekerjaan: "dll",
			Alasan:    "dll",
		},
		Profile{
			Code:      "GLNG012ONL003",
			Nama:      "Windra Rasyad",
			Alamat:    "Kuningan",
			Pekerjaan: "dll",
			Alasan:    "dll",
		},
		Profile{
			Code:      "GLNG012ONL004",
			Nama:      "Andre Philo",
			Alamat:    "Cawang",
			Pekerjaan: "dll",
			Alasan:    "dll",
		},
		Profile{
			Code:      "GLNG012ONL005",
			Nama:      "Reydika Ilham",
			Alamat:    "Menteng",
			Pekerjaan: "dll",
			Alasan:    "dll",
		},
	}
	for _, val := range data {
		if args[0] == val.Code {
			printBiodata(profile)
		}
	}
}
