package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	var argsRaw = os.Args
	var num = argsRaw[1]

		number, _ := strconv.Atoi(num)
	

	var dataSiswa = []Person{
		{Nama: "Dimas", Alamat: "Jl Melati", Pekerjaan: "Mahasiswa", Alasan: "Ingin mengetahui lebih dalam tentang pemrogaman"},
		{Nama: "Faqih", Alamat: "Jl perkutut", Pekerjaan: "Mahasiswa", Alasan: "Ingin menjadi programmer handal"},
		{Nama: "Alfirman", Alamat: "Jl Bulak", Pekerjaan: "Mahasiswa", Alasan: "Ingin jadi backend engineer professional"},
		{Nama: "Indra", Alamat: "Jl Pondok indah", Pekerjaan: "Mahasiswa", Alasan: "Ingin cari Modal buat cari kerja"},
		{Nama: "Fahrizal", Alamat: "Jl Ciledug", Pekerjaan: "PNS", Alasan: "Ingin memperdalam ilmu coding"},
		{Nama: "Rafli", Alamat: "Jl Apel", Pekerjaan: "Mahasiswa", Alasan: "Ingin dapetin skill baru"},
		{Nama: "Samsul", Alamat: "Jl Senopati", Pekerjaan: "Pengangguran", Alasan: "Ingin nyari kesibukan"},
		{Nama: "Eko", Alamat: "Jl Kenangan ", Pekerjaan: "Mahasiswa", Alasan: "Karena diajak teman"},
		{Nama: "Iwan", Alamat: "Jl Sudirman", Pekerjaan: "Mahasiswa", Alasan: "Karen disuruh orang tua"},
		{Nama: "Tedi", Alamat: "Jl Kebangsaan", Pekerjaan: "Freelancer", Alasan: "Ingin Dapetin sertifikat"},
	}

	if number > len(dataSiswa) {
		showError()
	} else {
		showData(dataSiswa[number-1])
	}
}

func showData(siswa Person) {
	fmt.Println("Nama      : \t", siswa.Nama)
	fmt.Println("Alamat    : \t", siswa.Alamat)
	fmt.Println("Pekerjaan : \t", siswa.Pekerjaan)
	fmt.Println("Alasan    : \t", siswa.Alasan)
}

func showError() {
	fmt.Println("No Absen Tidak Ditemukan")
}