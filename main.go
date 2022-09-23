package main

import (
	"dts-h8-assignment-1/pkg"
	"fmt"
	"os"
)

var karyawan = []pkg.Karyawan{
	{
		Nama:                   "Mbah Kekong",
		Alamat:                 "Bikini Bottom",
		Pekerjaan:              "Peramal Cuaca, Dukun 4.0",
		AlasanPilihKelasGolang: "Cari jodoh"},
	{
		Nama:                   "Kang Halusin Nasi",
		Alamat:                 "Jl. Terus keliling komplek",
		Pekerjaan:              "Penjual bubur, Atlet kentongan",
		AlasanPilihKelasGolang: "Cari ilmu dan teman - teman baru",
	},
	{
		Nama:                   "Oyen",
		Alamat:                 "Jl. Kasur rusak",
		Pekerjaan:              "Tidur, main, makan",
		AlasanPilihKelasGolang: "Mengisi kegabutan",
	},
	{
		Nama:                   "Jenny Wong",
		Alamat:                 "Jl. Buntu",
		Pekerjaan:              "UX Designer",
		AlasanPilihKelasGolang: "Cari ilmu dan meningkatkan skill",
	},
	{
		Nama:                   "Annish Vargarant",
		Alamat:                 "Jl. Lah sebelum berlari",
		Pekerjaan:              "Content Writer",
		AlasanPilihKelasGolang: "Mencari partner untuk membangun Start-Up",
	},
}

func main() {
	a := os.Args
	fmt.Println(a)
}
