package main

import (
	"fmt"
)

	//Deklarasi variabel
	var (
		Username = "Nurul"
		Password = "2406361284"
		JenisKelamin = "Perempuan"
		ProgramStudi = "Produksi Media "
		Angkatan = "2024"
		Menu int 
		Histori []string
	)

//Struct untuk Buku
type Buku struct {
	NamaBuku string
	Jumlah int 
}

//Fungsi untuk melihat informasi pengguna
func LihatInformasiPenggunaProgram() {
	fmt.Println("=== Informasi Pengguna Akun ===")
	fmt.Printf("Username: %s\n", Username)
	fmt.Printf("Password: %s\n", Password)
	fmt.Printf("Jenis Kelamin: %s\n", JenisKelamin)
	fmt.Printf("Program Studi: %s\n", ProgramStudi)
	fmt.Printf("Angkatan: %s\n", Angkatan)
}

//Fungsi untuk melihat daftar buku
func LihatDaftarBuku(daftarBuku []Buku) {
	fmt.Println("Daftar Buku:")
	for i, buku := range daftarBuku {
		fmt.Printf("%d. Nama Buku: %s\n   Jumlah: %d\n", i+1, buku.NamaBuku, buku.Jumlah)
	}
}

//Fungsi untuk tambah daftar buku 
func TambahDaftarBuku(daftarBuku *[]Buku) {
	var namaBuku string
	var jumlah int 

	fmt.Print("Masukan Nama Buku: ")
	fmt.Scan(&namaBuku)
	
	fmt.Print("Masukan Jumlah Buku: ")
	fmt.Scan(&jumlah)

	//Menambahan buku baru kedalam daftar buku
	*daftarBuku = append(*daftarBuku, Buku{NamaBuku: namaBuku, Jumlah: jumlah})

	fmt.Println("Buku Baru berhasil ditambahkan!")
}

// Fungsi untuk tambah pinjaman buku 
func TambahPeminjamanBuku(daftarBuku *[]Buku) {
	var nomorBuku int
	var jumlahPeminjaman int

	// Tampilkan daftar buku
	fmt.Println("Pinjam Buku")
	for i, buku := range *daftarBuku {
		fmt.Printf("%d. %s (%d)\n", i+1, buku.NamaBuku, buku.Jumlah)
	}

	fmt.Print("Masukkan Nomor Pinjaman Buku: ")
	fmt.Scan(&nomorBuku)

	// Validasi nomor buku
	if nomorBuku < 1 || nomorBuku > len(*daftarBuku) {
		fmt.Println("Nomor buku tidak valid.")
		return
	}

	fmt.Print("Masukkan Jumlah Pinjaman: ")
	fmt.Scan(&jumlahPeminjaman)

	if jumlahPeminjaman <= 0 {
		fmt.Println("Jumlah peminjaman harus lebih dari 0")
		return
	}

	// Mendapatkan indeks buku yang dipilih
	bukuIndex := nomorBuku - 1
	buku := (*daftarBuku)[bukuIndex]

	if jumlahPeminjaman > buku.Jumlah {
		fmt.Println("Jumlah peminjaman melebihi jumlah buku yang tersedia.")
		return
	}

	// Mengurangi jumlah buku
	(*daftarBuku)[bukuIndex].Jumlah -= jumlahPeminjaman
	Histori = append(Histori, fmt.Sprintf("Dipinjam: %s, Jumlah: %d", buku.NamaBuku, jumlahPeminjaman))
	fmt.Printf("Buku '%s' sebanyak %d berhasil dipinjam.\n", buku.NamaBuku, jumlahPeminjaman)
}

//Fungsi untuk melihat histori peminjaman buku
func HistoriPeminjamanBuku() {
	fmt.Println("=== Histori Peminjaman Buku ===")
	if len(Histori) == 0 {
		fmt.Println("Belum ada histori peminjaman")
		return
	}
	fmt.Println("Histori terbaru: ")
	for _, entry := range Histori {
		fmt.Println(entry)
	}
}


func main() {
	//Variabel Input
	var InputUsername, InputPassword string

	//Daftar buku
	daftarBuku := []Buku{
		Buku{
			NamaBuku: "Pemprograman",
			Jumlah: 10,
		},
		Buku{
			NamaBuku: "Film",
			Jumlah: 5, 
		},
		Buku{
			NamaBuku: "Printing",
			Jumlah: 20,
		},
	}

	fmt.Println(daftarBuku[0].Jumlah)

	//Salam Pembuka 
	fmt.Println("=== Selamat Datang di Perpustakaan Vokasi Ui ===")

	//Input Username 
	fmt.Print("Silahkan Input Username: ")
	fmt.Scanf("%s\n", &InputUsername)

	//Input Password
	fmt.Print("Silahkan Input Password: ")
	fmt.Scanf("%s\n", &InputPassword)

	//Vaidasi Username & Password, Jika salah keduanya program akan berhenti 
	if InputUsername == Username && InputPassword == Password {
		fmt.Println("=== Login Sukses!!! ===")
		for {
			fmt.Println("=== Menu Program ===")
			fmt.Println("1. Lihat Informasi Pengguna Program")
			fmt.Println("2. Lihat Daftar Buku")
			fmt.Println("3. Tambah Daftar Buku")
			fmt.Println("4. Tambah Peminjaman Buku")
			fmt.Println("5. Histori Peminjaman Buku")
			fmt.Println("6. Keluar Dari Program")

			var Pilihan int 
			fmt.Print("Input Menu yang Anda Inginkan: ")
			fmt.Scan(&Pilihan)

			switch Pilihan{
			case 1: 
				LihatInformasiPenggunaProgram()
			case 2: 
				LihatDaftarBuku(daftarBuku)
			case 3: 
				TambahDaftarBuku(&daftarBuku)
			case 4: 
				TambahPeminjamanBuku(&daftarBuku)
			case 5: 
				HistoriPeminjamanBuku()
			case 6: 
				fmt.Println("=== Terima kasih telah menggunakan program peminjaman buku. Sampai jumpa! ===")
				return
			default: 
				fmt.Println("Pilihan tidak ada.")

			}
		}
	} else {
		fmt.Println("Username atau Password Anda Salah.")
		return
	}
} 
