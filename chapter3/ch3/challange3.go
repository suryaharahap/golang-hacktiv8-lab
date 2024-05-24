package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var listFriend = []Biodata{
	{Nama: "Surya", Alamat: "Jakarta", Pekerjaan: "Software Engineer", Alasan: "Ingin belajar Golang"},
	{Nama: "Dani", Alamat: "Bandung", Pekerjaan: "Data Scientist", Alasan: "Ingin meningkatkan skill programming - golang"},
	{Nama: "Airell", Alamat: "Surabaya", Pekerjaan: "Web Developer", Alasan: "Ingin belajar bahasa pemrograman baru"},
	{Nama: "Hunis", Alamat: "Medan", Pekerjaan: "Mobile Developer", Alasan: "Ingin switch karir ke bidang backend engineer"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silahkan masukkan nomor absen teman anda.")
		return
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		return
	}

	index := absen - 1
	if index < 0 || index >= len(listFriend) {
		fmt.Println("Data teman dengan nomor absen", absen, "tidak ditemukan/tidak ada.")
		return
	}

	friend := listFriend[index]
	fmt.Println("Nama:", friend.Nama)
	fmt.Println("Alamat:", friend.Alamat)
	fmt.Println("Pekerjaan:", friend.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", friend.Alasan)
}
