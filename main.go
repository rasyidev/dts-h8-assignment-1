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
func findFriendByAbsentNumber(noAbsen int) (friend pkg.Friend, err error) {
	if noAbsen > len(friends) {
		err = errors.New("jumlah teman di luar jangkauan")
		return friend, err
	}

	friend = friends[noAbsen-1]
	return friend, nil
}

func printFriendName(friend pkg.Friend) {
	fmt.Println("Nama\t\t\t\t: ", friend.Nama)
	fmt.Println("Alamat\t\t\t\t: ", friend.Alamat)
	fmt.Println("Pekerjaan\t\t\t: ", friend.Pekerjaan)
	fmt.Println("Alasan Pilih Kelas Golang\t: ", friend.AlasanPilihKelasGolang)
}

func parseArgs(args []string) (noAbsen int, err error) {
	if len(args) < 2 {
		err = errors.New("Tidak ada argument yang dikirim")
		return -1, err
	} else if len(args) > 2 {
		err = errors.New("Hanya menerima satu argument. Tidak boleh lebih dari satu")
		return -1, err
	} else {
		noAbsen, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err.Error())
		}

		return noAbsen, nil
	}
}

func main() {
	noAbsen, err := parseArgs(os.Args)
	if err != nil {
		panic(err.Error())
	}

	friend, err := findFriendByAbsentNumber(noAbsen)
	if err != nil {
		panic(err.Error())
	}

	printFriendName(friend)

}
