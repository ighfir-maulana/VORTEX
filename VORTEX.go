package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tanaman struct {
	ID             int
	Nama           string
	Jenis          string
	JumlahTanaman  int
	KondisiTanaman string
}

type HasilPanen struct {
	IDPanen   int
	IDTanaman int
	Jumlah    float64
	Kualitas  string
}

type PenggunaanBahan struct {
	IDPenggunaan int
	Tanggal      string
	PupukCair    float64
	Pestisida    float64
}

const MAX_TANAMAN = 100
const MAX_HASIL_PANEN = 100
const MAX_PENGGUNAAN_BAHAN = 100

var riwayatPenggunaanBahan [MAX_PENGGUNAAN_BAHAN]PenggunaanBahan
var jumlahPenggunaanSaatIni int
var nextPenggunaanID = 1

var kebun [MAX_TANAMAN]Tanaman
var jumlahTanamanSaatIni int
var nextTanamanID = 1

var hasilPanenData [MAX_HASIL_PANEN]HasilPanen
var jumlahHasilPanenSaatIni int
var nextPanenID = 1

func bacaInputString(inputSTR string) string {
	fmt.Print(inputSTR)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func bacaInputInteger(input string) int {
	for {
		inputStr := bacaInputString(input)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid. Harap masukkan angka bulat.")
	}
}

func bacaInputFloat(input string) float64 {
	for {
		inputStr := bacaInputString(input)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid. Harap masukkan angka desimal.")
	}
}

func main() {
	var pilihanUtama int
	for {
		fmt.Println("\nSelamat Datang di VORTEX: Vertical Optimized Resource Tracking and EXecution")
		tampilkanMenuUtama()
		pilihanUtama = bacaInputInteger("Masukkan pilihan Anda: ")

		switch pilihanUtama {
		case 1:
			MenuKelolaTanaman()
		case 2:
			MenuKelolaPanen()
		case 3:
			catatPenggunaanBahan()
		case 4:
			rekomendasiTindakan()
		case 5:
			MenuTampilkanData()
		case 0:
			fmt.Println("Terima kasih telah menggunakan VORTEX. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		bacaInputString("\nTekan ENTER untuk melanjutkan ke menu utama...")
	}
}
func tampilkanMenuUtama() {
	fmt.Println("\n--- Menu Utama VORTEX ---")
	fmt.Println("1. Kelola Tanaman")
	fmt.Println("2. Kelola Panen")
	fmt.Println("3. Catat Penggunaan Bahan")
	fmt.Println("4. Rekomendasi Tindakan")
	fmt.Println("5. Tampilkan Data")
	fmt.Println("0. Keluar")
	fmt.Println("--------------------------")
}

func tampilkanMenuKelolaTanaman() {
	fmt.Println("\n*** Menu Kelola Tanaman ***")
	fmt.Println("1. Tambah Tanaman")
	fmt.Println("2. Ubah Tanaman")
	fmt.Println("3. Hapus Tanaman")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Println("***************************")
}

func MenuKelolaTanaman() {
	var pilihan int
	for {
		tampilkanMenuKelolaTanaman()
		pilihan = bacaInputInteger("Masukkan pilihan Anda: ")
		switch pilihan {
		case 1:
			tambahTanaman()
		case 2:
			ubahTanaman()
		case 3:
			hapusTanaman()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		bacaInputString("\nTekan ENTER untuk melanjutkan...")
	}
}

func tambahTanaman() {
	if jumlahTanamanSaatIni >= MAX_TANAMAN {
		fmt.Println("Kebun sudah penuh, tidak dapat menambah tanaman baru.")
		return
	}

	fmt.Println("\n--- Tambah Tanaman Baru ---")
	tanamanBaru := Tanaman{}
	tanamanBaru.ID = nextTanamanID
	nextTanamanID++

	tanamanBaru.Nama = strings.Title(bacaInputString("Nama Tanaman: "))
	tanamanBaru.Jenis = strings.Title(bacaInputString("Jenis/Varietas: "))
	tanamanBaru.JumlahTanaman = bacaInputInteger("Jumlah Tanaman (per modul/unit): ")
	tanamanBaru.KondisiTanaman = strings.Title(bacaInputString("Kondisi Tanaman Saat Ini (Sehat/Layu/Terinfeksi/Berbunga/Berbuah): "))

	kebun[jumlahTanamanSaatIni] = tanamanBaru
	jumlahTanamanSaatIni++
	fmt.Println("Tanaman berhasil ditambahkan dengan ID:", tanamanBaru.ID)
}

func ubahTanaman() {
	if jumlahTanamanSaatIni == 0 {
		fmt.Println("Tidak ada data tanaman untuk diubah.")
		return
	}
	fmt.Println("\n--- Ubah Data Tanaman ---")
	idDicari := bacaInputInteger("Masukkan ID Tanaman yang ingin diubah: ")

	SelectionSortTanaman(&kebun, jumlahTanamanSaatIni, "ID", "ascending")
	indeksDitemukan := binarySearchTanamanByID(idDicari)

	if indeksDitemukan != -1 {
		fmt.Println("Data Tanaman Ditemukan. Masukkan data baru (kosongkan jika tidak ingin diubah):")

		namaBaru := bacaInputString(fmt.Sprintf("Nama Tanaman (Sebelumnya: %s): ", kebun[indeksDitemukan].Nama))
		if namaBaru != "" {
			kebun[indeksDitemukan].Nama = strings.Title((namaBaru))
		}

		jenisBaru := bacaInputString(fmt.Sprintf("Jenis/Varietas (Sebelumnya: %s): ", kebun[indeksDitemukan].Jenis))
		if jenisBaru != "" {
			kebun[indeksDitemukan].Jenis = strings.Title((jenisBaru))
		}

		fmt.Print(fmt.Sprintf("Jumlah Tanaman (Sebelumnya: %d): ", kebun[indeksDitemukan].JumlahTanaman))
		inputJumlah := bacaInputString("")
		if inputJumlah != "" {
			val, err := strconv.Atoi(inputJumlah)
			if err == nil {
				kebun[indeksDitemukan].JumlahTanaman = val
			} else {
				fmt.Println("Input jumlah tidak valid, data jumlah tidak diubah.")
			}
		}

		kondisiBaru := bacaInputString(fmt.Sprintf("Kondisi Tanaman (Sebelumnya: %s): ", kebun[indeksDitemukan].KondisiTanaman))
		if kondisiBaru != "" {
			kebun[indeksDitemukan].KondisiTanaman = strings.Title(kondisiBaru)
		}

		fmt.Println("Data tanaman dengan ID", idDicari, "berhasil diubah.")
	} else {
		fmt.Println("Tanaman dengan ID", idDicari, "tidak ditemukan.")
	}
}

func hapusTanaman() {
	if jumlahTanamanSaatIni == 0 {
		fmt.Println("Tidak ada data tanaman untuk dihapus.")
		return
	}

	fmt.Println("\n--- Hapus Data Tanaman ---")
	idDicari := bacaInputInteger("Masukkan ID Tanaman yang ingin dihapus: ")

	SelectionSortTanaman(&kebun, jumlahTanamanSaatIni, "ID", "ascending")
	indeksDitemukan := binarySearchTanamanByID(idDicari)

	if indeksDitemukan != -1 {
		konfirmasi := strings.ToLower(bacaInputString(fmt.Sprintf("Anda yakin ingin menghapus tanaman '%s', ID: %d dengan semua data hasil panen terkait? (ya/tidak): ", kebun[indeksDitemukan].Nama, idDicari)))
		if konfirmasi == "ya" {
			for i := indeksDitemukan; i < jumlahTanamanSaatIni-1; i++ {
				kebun[i] = kebun[i+1]
			}
			jumlahTanamanSaatIni--
			kebun[jumlahTanamanSaatIni] = Tanaman{}
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	} else {
		fmt.Println("Tanaman dengan ID", idDicari, "tidak ditemukan.")
	}
}
func tampilkanMenuKelolaPanen() {
	fmt.Println("\n--- Menu Kelola Panen ---")
	fmt.Println("1. Tambah Hasil Panen")
	fmt.Println("2. Ubah Hasil Panen")
	fmt.Println("3. Hapus Hasil Panen")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Println("-------------------------")
}

func MenuKelolaPanen() {
	var pilihan int
	for {
		tampilkanMenuKelolaPanen()
		pilihan = bacaInputInteger("Masukkan pilihan Anda: ")
		switch pilihan {
		case 1:
			tambahHasilPanen()
		case 2:
			ubahHasilPanen()
		case 3:
			hapusHasilPanen()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		bacaInputString("\nTekan ENTER untuk melanjutkan...")
	}
}

func tambahHasilPanen() {
	if jumlahHasilPanenSaatIni >= MAX_HASIL_PANEN {
		fmt.Println("Kapasitas penyimpanan hasil panen penuh.")
		return
	}

	fmt.Println("\n--- Tambah Data Hasil Panen ---")
	hasilPanenBaru := HasilPanen{}
	hasilPanenBaru.IDPanen = nextPanenID
	nextPanenID++

	tanamanDitemukan := false
	for !tanamanDitemukan {
		idTanaman := bacaInputInteger("ID Tanaman (yang dipanen): ")
		if SequentialSearchTanamanID(idTanaman) != -1 {
			hasilPanenBaru.IDTanaman = idTanaman
			tanamanDitemukan = true
		} else {
			fmt.Println("ID Tanaman tidak ditemukan. Harap masukkan ID Tanaman yang valid.")
		}
	}

	hasilPanenBaru.Jumlah = bacaInputFloat("Jumlah Hasil Panen (kg): ")
	hasilPanenBaru.Kualitas = bacaInputString("Kualitas (Baik/Sedang/Kurang Baik): ")

	hasilPanenData[jumlahHasilPanenSaatIni] = hasilPanenBaru
	jumlahHasilPanenSaatIni++
	fmt.Println("Data hasil panen berhasil ditambahkan dengan ID Panen:", hasilPanenBaru.IDPanen)
}

func ubahHasilPanen() {
	if jumlahHasilPanenSaatIni == 0 {
		fmt.Println("Tidak ada data hasil panen untuk diubah.")
		return
	}
	fmt.Println("\n--- Ubah Data Hasil Panen ---")
	idPanenDicari := bacaInputInteger("Masukkan ID Panen yang ingin diubah: ")

	indeksDitemukan := SequentialSearchHasilPanenByID(idPanenDicari)

	if indeksDitemukan != -1 {
		fmt.Println("Data Hasil Panen Ditemukan. Masukkan data baru (kosongkan jika tidak ingin diubah):")

		// ID Tanaman tidak diubah karena sudah terikat dengan panen tersebut
		fmt.Printf("ID Tanaman (Tidak dapat diubah): %d\n", hasilPanenData[indeksDitemukan].IDTanaman)

		fmt.Print(fmt.Sprintf("Jumlah Hasil Panen (Sebelumnya: %.2f kg): ", hasilPanenData[indeksDitemukan].Jumlah))
		inputJumlah := bacaInputString("")
		if inputJumlah != "" {
			val, err := strconv.ParseFloat(inputJumlah, 64)
			if err == nil {
				hasilPanenData[indeksDitemukan].Jumlah = val
			} else {
				fmt.Println("Input jumlah tidak valid, data jumlah tidak diubah.")
			}
		}

		kualitasBaru := bacaInputString(fmt.Sprintf("Kualitas (Sebelumnya: %s): ", hasilPanenData[indeksDitemukan].Kualitas))
		if kualitasBaru != "" {
			hasilPanenData[indeksDitemukan].Kualitas = strings.Title(kualitasBaru)
		}

		fmt.Println("Data hasil panen dengan ID Panen", idPanenDicari, "berhasil diubah.")
	} else {
		fmt.Println("Hasil panen dengan ID Panen", idPanenDicari, "tidak ditemukan.")
	}
}

func hapusHasilPanen() {
	if jumlahHasilPanenSaatIni == 0 {
		fmt.Println("Tidak ada data hasil panen untuk dihapus.")
		return
	}

	fmt.Println("\n--- Hapus Data Hasil Panen ---")
	idPanenDicari := bacaInputInteger("Masukkan ID Panen yang ingin dihapus: ")

	indeksDitemukan := SequentialSearchHasilPanenByID(idPanenDicari)

	if indeksDitemukan != -1 {
		konfirmasi := strings.ToLower(bacaInputString(fmt.Sprintf("Anda yakin ingin menghapus hasil panen ID: %d, Jumlah: %.2f kg? (ya/tidak): ", hasilPanenData[indeksDitemukan].IDPanen, hasilPanenData[indeksDitemukan].Jumlah)))
		if konfirmasi == "ya" {
			for i := indeksDitemukan; i < jumlahHasilPanenSaatIni-1; i++ {
				hasilPanenData[i] = hasilPanenData[i+1]
			}
			jumlahHasilPanenSaatIni--
			hasilPanenData[jumlahHasilPanenSaatIni] = HasilPanen{}

			fmt.Println("Hasil panen dengan ID Panen", idPanenDicari, "berhasil dihapus.")
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	} else {
		fmt.Println("Hasil panen dengan ID Panen", idPanenDicari, "tidak ditemukan.")
	}
}

func catatPenggunaanBahan() {
	if jumlahPenggunaanSaatIni >= MAX_PENGGUNAAN_BAHAN {
		fmt.Println("Kapasitas penyimpanan catatan penggunaan bahan penuh.")
		return
	}

	fmt.Println("\n--- Catat Penggunaan Bahan ---")
	penggunaanBaru := PenggunaanBahan{}
	penggunaanBaru.IDPenggunaan = nextPenggunaanID
	nextPenggunaanID++

	penggunaanBaru.Tanggal = bacaInputString("Tanggal Penggunaan (YYYY-MM-DD): ")
	penggunaanBaru.PupukCair = bacaInputFloat("Jumlah Pupuk Cair yang Digunakan (liter): ")
	penggunaanBaru.Pestisida = bacaInputFloat("Jumlah Pestisida yang Digunakan (liter): ")

	riwayatPenggunaanBahan[jumlahPenggunaanSaatIni] = penggunaanBaru
	jumlahPenggunaanSaatIni++
	fmt.Println("Catatan penggunaan bahan berhasil ditambahkan dengan ID:", penggunaanBaru.IDPenggunaan)
}

func rekomendasiTindakan() {
	if jumlahTanamanSaatIni == 0 {
		fmt.Println("Tidak ada data tanaman untuk mendapatkan rekomendasi tindakan.")
		return
	}
	fmt.Println("\n--- Rekomendasi Tindakan Berdasarkan Kondisi Tanaman ---")
	idDicari := bacaInputInteger("Masukkan ID Tanaman yang ingin dapat rekomendasi: ")

	indeksDitemukan := SequentialSearchTanamanID(idDicari)

	if indeksDitemukan != -1 {
		tanaman := kebun[indeksDitemukan]
		fmt.Println("\nRekomendasi Tindakan untuk Tanaman:", tanaman.Nama, "(ID:", tanaman.ID, ")", "(Kondisi: )", tanaman.KondisiTanaman)
		switch strings.ToLower(tanaman.KondisiTanaman) {
		case "layu":
			fmt.Println("- Periksa ketersediaan air dan kelembaban tanah.")
			fmt.Println("- Periksa apakah ada tanda-tanda kekurangan nutrisi.")
		case "terinfeksi":
			fmt.Println("- Identifikasi jenis infeksi (hama/penyakit).")
			fmt.Println("- Cari informasi mengenai cara pengendalian yang tepat (organik/kimiawi).")
			fmt.Println("- Pertimbangkan untuk mengisolasi tanaman yang terinfeksi untuk mencegah penyebaran.")
		case "berbunga":
			fmt.Println("- Pastikan kondisi lingkungan optimal untuk penyerbukan.")
			fmt.Println("- Berikan nutrisi yang cukup untuk mendukung pembentukan bunga dan buah.")
		case "berbuah":
			fmt.Println("- Monitor perkembangan buah secara berkala.")
			fmt.Println("- Pastikan nutrisi yang cukup untuk pembesaran buah.")
			fmt.Println("- Pertimbangkan penyangga jika buah terlalu berat.")
		case "sehat":
			fmt.Println("- Tanaman dalam kondisi sehat. Terus pantau dan jaga kondisi optimal.")
			fmt.Println("- Lakukan pemeliharaan rutin seperti pemangkasan jika diperlukan.")
		default:
			fmt.Println("- Kondisi tanaman:", tanaman.KondisiTanaman)
			fmt.Println("- Belum ada rekomendasi khusus untuk kondisi ini.")
		}
	} else {
		fmt.Println("Tanaman dengan ID", idDicari, "tidak ditemukan.")
	}
}

func tampilkanMenuTampilkanData() {
	fmt.Println("\n--- Menu Tampilkan Data ---")
	fmt.Println("1. Tampilkan & Urutkan Tanaman")
	fmt.Println("2. Tampilkan & Urutkan Hasil Panen")
	fmt.Println("3. Tampilkan Riwayat Penggunaan Bahan")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Println("-----------------------------")
}

func MenuTampilkanData() {
	var pilihan int
	for {
		tampilkanMenuTampilkanData()
		pilihan = bacaInputInteger("Masukkan pilihan Anda: ")
		switch pilihan {
		case 1:
			tampilkanDanUrutkanTanaman()
		case 2:
			tampilkanDanUrutkanHasilPanen()
		case 3:
			tampilkanRiwayatPenggunaanBahan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		bacaInputString("\nTekan ENTER untuk melanjutkan...")
	}
}

func tampilkanDanUrutkanTanaman() {
	if jumlahTanamanSaatIni == 0 {
		fmt.Println("Tidak ada data tanaman untuk ditampilkan.")
		return
	}
	fmt.Println("\n--- Tampilkan & Urutkan Tanaman ---")
	fmt.Println("Berdasarkan: (1) Nama (2) ID")
	pilihanUSER := bacaInputInteger("Pilihan Anda: ")
	fmt.Println("\nUrutan: (1) Ascending (2) Descending")
	pilihanASCDESC := bacaInputInteger("Pilihan Anda: ")

	var kriteriaStr string
	var urutanStr string
	cekInput := true

	if pilihanUSER == 1 {
		kriteriaStr = "Nama"
	} else if pilihanUSER == 2 {
		kriteriaStr = "ID"
	} else {
		fmt.Println("Kriteria tidak valid.")
		cekInput = false
	}

	if cekInput {
		if pilihanASCDESC == 1 {
			urutanStr = "ascending"
		} else if pilihanASCDESC == 2 {
			urutanStr = "descending"
		} else {
			fmt.Println("Urutan tidak valid.")
			cekInput = false
		}
	}

	if cekInput {
		SelectionSortTanaman(&kebun, jumlahTanamanSaatIni, kriteriaStr, urutanStr)
		fmt.Println("\nTanaman berhasil diurutkan.")
		tampilkanTanaman(kebun, jumlahTanamanSaatIni)
	} else {
		fmt.Println("Input tidak valid untuk pengurutan. Pembatalan.")
	}
}

func SelectionSortTanaman(kebun *[MAX_TANAMAN]Tanaman, jumlahTanamanSaatIni int, berdasarkan string, urutan string) {
	for i := 0; i < jumlahTanamanSaatIni-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahTanamanSaatIni; j++ {
			var TukarNilai bool

			switch berdasarkan {
			case "Nama":
				if urutan == "ascending" {
					TukarNilai = kebun[j].Nama < kebun[minIdx].Nama
				} else {
					TukarNilai = kebun[j].Nama > kebun[minIdx].Nama
				}
			case "ID":
				if urutan == "ascending" {
					TukarNilai = kebun[j].ID < kebun[minIdx].ID
				} else {
					TukarNilai = kebun[j].ID > kebun[minIdx].ID
				}
			}

			if TukarNilai {
				minIdx = j
			}
		}
		kebun[i], kebun[minIdx] = kebun[minIdx], kebun[i]
	}
}

func tampilkanTanaman(kebun [MAX_TANAMAN]Tanaman, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Tidak ada data tanaman.")
		return
	}
	fmt.Println("\nDaftar Tanaman:")
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-15s %-10s %-15s\n", "ID", "Nama", "Jenis", "Jumlah", "Kondisi")
	fmt.Println("---------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-5d %-15s %-15s %-10d %-15s\n", kebun[i].ID, kebun[i].Nama, kebun[i].Jenis, kebun[i].JumlahTanaman, kebun[i].KondisiTanaman)
	}
	fmt.Println("---------------------------------------------------------------------------------")
}

func tampilkanDanUrutkanHasilPanen() {
	if jumlahHasilPanenSaatIni == 0 {
		fmt.Println("Tidak ada data hasil panen untuk ditampilkan.")
		return
	}

	fmt.Println("\n--- Tampilkan & Urutkan Hasil Panen ---")
	fmt.Println("Berdasarkan: (1) Jumlah (2) ID Tanaman (3) ID Panen")
	kriteriaPilihan := bacaInputInteger("Pilihan Anda: ")
	fmt.Println("Urutan: (1) Ascending (2) Descending")
	urutanPilihan := bacaInputInteger("Pilihan Anda: ")

	var kriteriaStr string
	var urutanStr string
	cekInput := true

	if kriteriaPilihan == 1 {
		kriteriaStr = "Jumlah"
	} else if kriteriaPilihan == 2 {
		kriteriaStr = "IDTanaman"
	} else if kriteriaPilihan == 3 {
		kriteriaStr = "IDPanen"
	} else {
		fmt.Println("Kriteria tidak valid.")
		cekInput = false
	}

	if cekInput {
		if urutanPilihan == 1 {
			urutanStr = "ascending"
		} else if urutanPilihan == 2 {
			urutanStr = "descending"
		} else {
			fmt.Println("Urutan tidak valid.")
			cekInput = false
		}
	}

	if cekInput {
		HasilPanenInsertionSort(&hasilPanenData, jumlahHasilPanenSaatIni, kriteriaStr, urutanStr)
		fmt.Println("\nHasil panen berhasil diurutkan.")
		tampilkanHasilPanen(hasilPanenData, jumlahHasilPanenSaatIni)
	} else {
		fmt.Println("Input tidak valid untuk pengurutan. Pembatalan.")
	}
}

func HasilPanenInsertionSort(hasil *[MAX_HASIL_PANEN]HasilPanen, jumlahHasilPanenSaatIni int, berdasarkan string, urutan string) {
	for i := 1; i < jumlahHasilPanenSaatIni; i++ {
		temp := hasil[i]
		j := i - 1

		for j >= 0 && ((berdasarkan == "Jumlah" && ((urutan == "ascending" && hasil[j].Jumlah > temp.Jumlah) || (urutan == "descending" && hasil[j].Jumlah < temp.Jumlah))) ||
			(berdasarkan == "IDTanaman" && ((urutan == "ascending" && hasil[j].IDTanaman > temp.IDTanaman) || (urutan == "descending" && hasil[j].IDTanaman < temp.IDTanaman))) ||
			(berdasarkan == "IDPanen" && ((urutan == "ascending" && hasil[j].IDPanen > temp.IDPanen) || (urutan == "descending" && hasil[j].IDPanen < temp.IDPanen)))) {
			hasil[j+1] = hasil[j]
			j--
		}
		hasil[j+1] = temp
	}
}

func tampilkanHasilPanen(hasil [MAX_HASIL_PANEN]HasilPanen, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Tidak ada data hasil panen.")
		return
	}
	fmt.Println("\nDaftar Hasil Panen:")
	fmt.Println("*******************************************************")
	fmt.Printf("%-10s %-10s %-10s %-15s\n", "ID Panen", "ID Tanaman", "Jumlah(kg)", "Kualitas")
	fmt.Println("*******************************************************")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-10d %-10d %-10.2f %-15s\n", hasil[i].IDPanen, hasil[i].IDTanaman, hasil[i].Jumlah, hasil[i].Kualitas)
	}
	fmt.Println("-------------------------------------------------------")
}

func tampilkanRiwayatPenggunaanBahan() {
	if jumlahPenggunaanSaatIni == 0 {
		fmt.Println("\nBelum ada catatan penggunaan bahan.")
		return
	}

	fmt.Println("\n--- Riwayat Penggunaan Bahan ---")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-15s %-15s\n", "ID", "Tanggal", "Pupuk Cair (L)", "Pestisida (L)")
	fmt.Println("-------------------------------------------------------")
	for _, penggunaan := range riwayatPenggunaanBahan[:jumlahPenggunaanSaatIni] {
		fmt.Printf("%-5d %-15s %-15.2f %-15.2f\n", penggunaan.IDPenggunaan, penggunaan.Tanggal, penggunaan.PupukCair, penggunaan.Pestisida)
	}
	fmt.Println("-------------------------------------------------------")
}

func SequentialSearchTanamanID(id int) int {
	for i := 0; i < jumlahTanamanSaatIni; i++ {
		if kebun[i].ID == id {
			return i
		}
	}
	return -1
}

func binarySearchTanamanByID(id int) int {
	low := 0
	high := jumlahTanamanSaatIni - 1

	for low <= high {
		mid := (low + high) / 2
		if kebun[mid].ID == id {
			return mid
		} else if kebun[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func SequentialSearchHasilPanenByID(id int) int {
	for i := 0; i < jumlahHasilPanenSaatIni; i++ {
		if hasilPanenData[i].IDPanen == id {
			return i
		}
	}
	return -1
}
