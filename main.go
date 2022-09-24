package main

import (
	"dts-h8-assignment-1/pkg"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var friends = []pkg.Friend{
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

// Untuk mencari data teman dengan no absen tertentu
func findFriendByAbsentNumber(noAbsen int) (pkg.Friend, error) {
	if noAbsen > len(friends) {
		return pkg.Friend{}, errors.New("Jumlah teman di luar jangkauan")
	}

	friend := friends[noAbsen-1]
	return friend, nil
}

func printFriendName(friend pkg.Friend) {
	fmt.Println("Nama\t\t\t\t: ", friend.Nama)
	fmt.Println("Alamat\t\t\t\t: ", friend.Alamat)
	fmt.Println("Pekerjaan\t\t\t: ", friend.Pekerjaan)
	fmt.Println("Alasan Pilih Kelas Golang\t: ", friend.AlasanPilihKelasGolang)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Argumen tidak boleh kosong")
	} else if len(args) == 2 {
		noAbsen, _ := strconv.Atoi(args[1])
		friend, err := findFriendByAbsentNumber(noAbsen)
		if err != nil {
			panic(err.Error())
		}

		printFriendName(friend)

	} else {
		fmt.Println("Argumen tidak boleh lebih dari satu")
	}

}
