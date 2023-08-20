package main

import (
	"fmt"
	"strconv"
)

func main() {
	DMahasiswa := make(map[string]string)
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
	var searchKey string
	fmt.Print("Masukkan Nim Mahasiswa :")
	fmt.Scanln(&searchKey)

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
			mid = mid
			found = true
		}
	}

	if !found {
		fmt.Println("NIM tidak ditemukan.")
	} else {
		DMahasiswa["Nama"] = Data[mid][2]
		DMahasiswa["Nim"] = Data[mid][1]
		DMahasiswa["Kelas"] = Data[mid][0]

		fmt.Println("====================================")
		fmt.Println("Nama :", DMahasiswa["Nama"])
		fmt.Println("NIM :", DMahasiswa["Nim"])
		fmt.Println("Kelas :", DMahasiswa["Kelas"])
		fmt.Println("====================================")

		var (
			Ntugas, Nkehadiran, Nfinal float32
			tugas, Jkehadiran          int16
			NAkhir                     string
		)
		Ntugas = 0
		count = 1
		Jtugas := 2
		for count <= Jtugas {
			fmt.Print("Masukkan Nilai Tugas Ke ", count, ":")
			fmt.Scanln(&tugas)
			Ntugas += float32(tugas)
			count++
		}
		Ntugas = Ntugas / float32(Jtugas)
		fmt.Print("Masukkan Jumlah Kehadiran :")
		fmt.Scanln(&Jkehadiran)
		Nkehadiran = (float32(Jkehadiran) * 100) / 16

		fmt.Print("Masukkan Nilai Final :")
		fmt.Scanln(&Nfinal)

		Nfinal = (Nfinal + Nkehadiran + Ntugas) / 3
		NAkhir = strconv.FormatFloat(float64(Nfinal), 'f', -1, 32)
		DMahasiswa["Nilai"] = NAkhir

		fmt.Println("Nama :", DMahasiswa["Nama"])
		fmt.Println("NIM :", DMahasiswa["Nim"])
		fmt.Println("Kelas :", DMahasiswa["Kelas"])
		fmt.Println("Nilai Akhir :", DMahasiswa["Nilai"])
	}

	fmt.Println("====================================")
	fmt.Println("=========|Program Berakhir|=========")
	fmt.Println("====================================")

}
