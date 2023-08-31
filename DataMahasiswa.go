package main

import (
	"fmt"
	"strings"
)

func DataMahasiswa(searchKey string) {
	DMahasiswa := make(map[string]string)
	for {
		var Data = [...][3]string{
			// {"kelas", "NIM", "Nama"}
			{"2D", "105841113822", "Abdullah Khaeruna Anwar"},   //0
			{"2D", "105841113922", "Muh. Irsyad Jafar"},         //1
			{"2D", "105841114022", "Alryadi Asmu'adzan"},        //2
			{"2D", "105841114122", "Muhammad Aditya Yudistira"}, //3
			{"2D", "105841114222", "Wiwin Fuad Sanjaya"},        //4
			{"2D", "105841114322", "Muh. Ayyub Hasrul"},         //5
			{"2D", "105841114422", "Muhammad Alif Syaffan"},     //6
			{"2D", "105841114522", "Muh. Imam Ma'ruf Musni"},    //7
			{"2D", "105841114622", "Muhammad Abdillah Zufar"},   //8
			{"2D", "105841114722", "Mutiara"},                   //9
			{"2D", "105841114822", "M. Fahmi Zulhidayat"},       //10
			{"2D", "105841114922", "Muh. Hasrul"},               //11
		}

		var rendah int = 0
		var tinggi int = len(Data) - 1
		var mid int

		found := false
		count := 0

		for rendah <= tinggi && !found {
			count++
			rendah1 := rendah
			tinggi1 := tinggi
			mid = (rendah + tinggi) / 2
			// Nim :=

			// strconv.Atoi(variabel) adalah perintah konversi data dari string ke int
			if Data[rendah1][1] == searchKey {
				mid = rendah1
				found = true
				break
			} else if Data[tinggi1][1] == searchKey {
				mid = tinggi1
				found = true
				break
			} else if Data[mid][1] < searchKey {
				rendah = mid + 1
				tinggi--
			} else if Data[mid][1] > searchKey {
				rendah++
				tinggi = mid - 1
			} else {
				found = true
			}
		}

		if !found {
			fmt.Println("NIM tidak ditemukan.")

			fmt.Print("Siahkan Masukkan Ulang NIM Yang Anda Cari :")
			fmt.Scanln(&searchKey)
		} else {
			DMahasiswa["Nama"] = Data[mid][2]
			DMahasiswa["Nim"] = Data[mid][1]
			DMahasiswa["Kelas"] = Data[mid][0]

			fmt.Println("Nama :", DMahasiswa["Nama"])
			fmt.Println("NIM :", DMahasiswa["Nim"])
			fmt.Println("Kelas :", DMahasiswa["Kelas"])

			break
		}
	}

}

func Len(Data [12][3]string) {
	panic("unimplemented")
}

func NilaiAkhir(Jtugas int) {
	var (
		Nkehadiran, Nfinal, count, Ntugas, tugas int
	)
	count = 1
	Ntugas = 0
	fmt.Print("Masukkan Jumlah Kehadiran :")
	fmt.Scanln(&Nkehadiran)

	Nkehadiran = (Nkehadiran * 20) / 16

	for count <= Jtugas {
		fmt.Print("Masukkan Nilai Tugas ", count, ":")
		fmt.Scanln(&tugas)

		Ntugas += tugas
		count++
	}
	Ntugas = (Ntugas * 30) / (Jtugas * 100)

	fmt.Print("Masukkan Nilai Final/Mid :")
	fmt.Scanln(&Nfinal)

	Nfinal = (Nfinal * 50) / 100

	Nfinal = Nkehadiran + Ntugas + Nfinal

	fmt.Println("Nilai Akhir :", Nfinal)
}

func main() {

	var (
		searchKey, lanjut string
		Jtugas            int
	)
	lanjut = "ya"

	fmt.Print("Masukkan Jumlah Tugas :")
	fmt.Scanln(&Jtugas)

	for strings.ToLower(lanjut) == strings.ToLower("Ya") {
		fmt.Print("Masukkan Nim Mahasiswa :")
		fmt.Scanln(&searchKey)

		DataMahasiswa(searchKey)

		NilaiAkhir(Jtugas)

		fmt.Print("Apakah Anda Ingin Melakukan Perhitungan Lagi? (Ya/Tidak) :")
		fmt.Scanln(&lanjut)

	}
	fmt.Println("==========================")
	fmt.Println("====|Program Berakhir|====")
	fmt.Println("==========================")
}
