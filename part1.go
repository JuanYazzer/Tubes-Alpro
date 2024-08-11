package main

import "fmt"

const NMAX = 1000

type uang [NMAX]int

type tenant struct {
	nama             string
	jumlahTransaksi  int
	uangPerTransaksi uang
	uangTenant       int
	uangAdmin        int
}
type Pengguna struct {
	username string
	password string
}

type tabTenant [NMAX]tenant

func main() {
	var tent tabTenant
	var jumlahTent int
	var choose int
	var choice string
	var username, password string
	var user Pengguna
	user = Pengguna{
		username: "admin",
		password: "admin",
	}
	loginBerhasil := false
	for !loginBerhasil && choice != "2" {
		fmt.Println("===========MENU===========")
		fmt.Println("1. Login")
		fmt.Println("2. Batal")
		fmt.Print("Pilih opsi (1 atau 2): ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)
			fmt.Println("==========================")
			// Cek username dan password
			if user.username == username && user.password == password {
				loginBerhasil = true
			}

			if loginBerhasil {
				fmt.Println("Login berhasil!")
				fmt.Println("Selamat Datang Admin Kantin!!!!")
				for choose != 5 {
					menuManipulasi(&choose)
					switch choose {
					case 1:
						menuTambah(&tent, &jumlahTent)
						urutTenantNama(&tent, &jumlahTent)
					case 2:
						menuHapus(&tent, &jumlahTent)
					case 3:
						menuUbah(&tent, &jumlahTent)
					case 4:
						urutTenantJumlahTransaksi(&tent, &jumlahTent)
						printTenantData(tent, jumlahTent)
					}
				}
				fmt.Println("Terima-kasih sudah menggunakan program kami!! :)")
			} else {
				fmt.Println("Login gagal! Username atau password salah. Silakan coba lagi.")
			}

		case "2":
			fmt.Println("Login dibatalkan. Terima kasih!")

		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1 atau 2.")
		}
	}
}

func menuManipulasi(choose *int) {
	fmt.Println("============================")
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("============================")
	fmt.Println("1. Tambah tenant")
	fmt.Println("2. Hapus tenant")
	fmt.Println("3. Ubah tenant")
	fmt.Println("4. Tampilkan tenant")
	fmt.Println("5. Exit")
	fmt.Println("============================")
	fmt.Print("Pilihan anda : ")
	fmt.Scan(choose)
	fmt.Println("============================")
}

func menuTambah(tent *tabTenant, p *int) {
	var choose int
	fmt.Println("Data apa yang ingin anda tambah?")
	fmt.Println("1.Tenant")
	fmt.Println("2.Transaksi tenant")
	fmt.Println("3.menu")
	fmt.Scan(&choose)
	if choose == 1 {
		tambahTenant(tent, p)
	} else if choose == 2 {
		printTenantData(*tent, *p)
		tambahTenantTransaksi(tent, p)
	} else if choose == 3 {
	} else {
		menuTambah(tent, p)
	}
}

func tambahTenant(tent *tabTenant, p *int) {
	var nama string
	fmt.Println("jika batal langsung masukkan 'selesai'")
	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)
	if *p < NMAX && nama != "selesai" {
		tent[*p].nama = nama
		isiDataTransaksi(&tent[*p])
		*p++
	}
	if *p == NMAX {
		fmt.Println("Kapasitas kantin penuh !!")
		fmt.Println("Kembali Ke menu")
	}
	if nama == "selesai" {
		fmt.Println("Kembali Ke menu")
	}
}

func tambahTenantTransaksi(tent *tabTenant, p *int) {
	var nama string
	var index int
	fmt.Println("Masukkan nama tenant yang ingin anda tambah data transaksinya")
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	index = cariTenantBinary(*tent, *p, nama)
	if index != -1 {
		printTenantDataSatu(tent[index])
		isiDataTransaksi(&tent[index])
		fmt.Println("Berikut data yang ada setelah di Tambah: ")
		printTenantData(*tent, *p)
	} else {
		fmt.Println("Data yang anda ingin tambah tidak dapat ditemukan")
	}
}

func isiDataTransaksi(tent *tenant) {
	var uang int
	fmt.Println("Masukkan bilangan negatif jika ingin selesai")
	fmt.Print("Masukkan Uang : ")
	fmt.Scan(&uang)
	for tent.jumlahTransaksi < NMAX && uang >= 0 {
		tent.uangPerTransaksi[tent.jumlahTransaksi] = uang
		tent.jumlahTransaksi++
		fmt.Println("Masukkan bilanngan negatif jika ingin selesai")
		fmt.Print("Masukkan uang : ")
		fmt.Scan(&uang)
	}
	tent.uangAdmin = totalUangAdmin(&tent.jumlahTransaksi, &tent.uangPerTransaksi)
	tent.uangTenant = totalUangTenant(&tent.jumlahTransaksi, &tent.uangPerTransaksi)
}

func totalUangAdmin(i *int, data *uang) int {
	var total int
	for j := 0; j < *i; j++ {
		total += data[j] * 25 / 100
	}
	return total
}

func totalUangTenant(i *int, data *uang) int {
	var total int
	for j := 0; j < *i; j++ {
		total += data[j] * 75 / 100
	}
	return total
}

func menuHapus(tent *tabTenant, j *int) {
	var choose int
	fmt.Println("Data apa yang ingin anda hapus?")
	fmt.Println("1.Tenant")
	fmt.Println("2.Transaksi tenant")
	fmt.Println("3.menu")
	fmt.Scan(&choose)
	if choose == 1 {
		printTenantData(*tent, *j)
		hapusTenantdata(tent, j)
	} else if choose == 2 {
		printTenantData(*tent, *j)
		hapusTenantTransaksi(tent, j)
	} else if choose == 3 {
	} else {
		menuHapus(tent, j)
	}
}

func hapusTenantdata(tent *tabTenant, j *int) {
	var nama string
	var index int
	fmt.Println("Masukkan nama tenant yang ingin anda hapus keseluruhan datanya")
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	index = cariTenantSequential(*tent, *j, nama)
	if index != -1 {
		fmt.Println("Menghapus tenant", tent[index].nama)
		for i := index; i < *j; i++ {
			tent[i] = tent[i+1]
		}
		*j = *j - 1
		fmt.Println("Berikut data yang ada setelah di Hapus: ")
		printTenantData(*tent, *j)
	} else {
		fmt.Println("Data yang anda ingin hapus tidak dapat ditemukan")
	}
}

func hapusTenantTransaksi(tent *tabTenant, p *int) {
	var nama string
	var index int
	fmt.Println("Masukkan nama tenant yang ingin anda hapus data transaksinya")
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	index = cariTenantBinary(*tent, *p, nama)
	if index != -1 {
		printTenantDataSatu(tent[index])
		hapusTenantTransaksiData(&tent[index])
		fmt.Println("Berikut data yang ada setelah di Hapus: ")
		printTenantData(*tent, *p)
	} else {
		fmt.Println("Data yang anda ingin hapus tidak dapat ditemukan")
	}
}

func hapusTenantTransaksiData(tent *tenant) {
	var i int
	fmt.Println("Masukkan data ke berapa yang ingin anda hapus?")
	fmt.Println("Masukkan", tent.jumlahTransaksi+1, "jika batal")
	fmt.Print("Masukkan angka :")
	fmt.Scan(&i)
	i = i - 1
	for i > tent.jumlahTransaksi || i < 0 {
		fmt.Println("Masukkan tidak valid masukkan lagi!")
		fmt.Print("Masukkan angka :")
		fmt.Scan(&i)
		i = i - 1
	}
	if i >= 0 && i < tent.jumlahTransaksi {
		for j := i; j < tent.jumlahTransaksi; j++ {
			tent.uangPerTransaksi[i] = tent.uangPerTransaksi[i+1]
		}
		tent.jumlahTransaksi--
		tent.uangAdmin = totalUangAdmin(&tent.jumlahTransaksi, &tent.uangPerTransaksi)
		tent.uangTenant = totalUangTenant(&tent.jumlahTransaksi, &tent.uangPerTransaksi)
	}
}

func menuUbah(tent *tabTenant, j *int) {
	var choose int
	fmt.Println("Data apa yang ingin anda ubah?")
	fmt.Println("1.Tenant")
	fmt.Println("2.Transaksi tenant")
	fmt.Println("3.menu")
	fmt.Scan(&choose)
	if choose == 1 {
		printTenantData(*tent, *j)
		ubahTenantNama(tent, j)
	} else if choose == 2 {
		printTenantData(*tent, *j)
		ubahTenantTransaksi(tent, j)
	} else if choose == 3 {
	} else {
		menuUbah(tent, j)
	}
}

func ubahTenantNama(tent *tabTenant, j *int) {
	var nama string
	var index int
	fmt.Println("Masukkan nama tenant yang ingin anda ubah namanya")
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	index = cariTenantSequential(*tent, *j, nama)
	if index != -1 {
		fmt.Println("Mengubah nama tenant", tent[index].nama)
		fmt.Print("Masukkan nama yang anda inginkan : ")
		fmt.Scan(&nama)
		tent[index].nama = nama
		fmt.Println("Berikut data yang ada setelah di ubah: ")
		printTenantData(*tent, *j)
	} else {
		fmt.Println("Data yang anda ingin ubah tidak dapat ditemukan")
	}
}

func ubahTenantTransaksi(tent *tabTenant, p *int) {
	var nama string
	var index int
	fmt.Println("Masukkan nama tenant yang ingin anda ubah data transaksinya")
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	index = cariTenantBinary(*tent, *p, nama)
	if index != -1 {
		printTenantDataSatu(tent[index])
		ubahTenantTransaksiData(&tent[index])
		fmt.Println("Berikut data yang ada setelah di ubah: ")
		printTenantData(*tent, *p)
	} else {
		fmt.Println("Data yang anda ingin ubah tidak dapat ditemukan")
	}
}

func ubahTenantTransaksiData(tent *tenant) {
	var i int
	var uang int
	fmt.Println("Masukkan data ke berapa yang ingin anda ubah?")
	fmt.Println("Masukkan", tent.jumlahTransaksi+1, "jika batal")
	fmt.Print("Masukkan angka :")
	fmt.Scan(&i)
	i = i - 1
	for i > tent.jumlahTransaksi || i < 0 {
		fmt.Println("Masukkan tidak valid masukkan lagi!")
		fmt.Print("Masukkan angka :")
		fmt.Scan(&i)
		i = i - 1
	}
	if i >= 0 && i < tent.jumlahTransaksi {
		fmt.Print("Masukkan uang yang diubah : ")
		fmt.Scan(&uang)
		for uang < 0 {
			fmt.Println("input tidak valid!")
			fmt.Print("Masukkan uang yang diubah : ")
			fmt.Scan(&uang)
		}
		tent.uangPerTransaksi[i] = uang
		tent.uangAdmin = totalUangAdmin(&tent.jumlahTransaksi, &tent.uangPerTransaksi)
		tent.uangTenant = totalUangTenant(&tent.jumlahTransaksi, &tent.uangPerTransaksi)
	}
}

func cariTenantSequential(tent tabTenant, j int, n string) int {
	var temp int
	temp = -1
	for i := 0; i < j; i++ {
		if tent[i].nama == n {
			temp = i
		}
	}
	return temp
}

func cariTenantBinary(tent tabTenant, j int, n string) int {
	var mid, left, right int
	mid = -1
	left = 0
	right = j

	for left <= right && mid == -1 {
		mid = left + (right-left)/2
		if tent[mid].nama < n {
			left = mid + 1
		} else if tent[mid].nama > n {
			right = mid - 1
		} else if tent[mid].nama == n {
			return mid
		}
	}

	return mid
}

func urutTenantJumlahTransaksi(tent *tabTenant, k *int) {
	var cek int
	var temp tenant
	var i, j int
	for i = 1; i < *k; i++ {
		cek = tent[i].jumlahTransaksi
		temp = tent[i]
		j = i - 1
		for j >= 0 && tent[j].jumlahTransaksi < cek {
			tent[j+1] = tent[j]
			j = j - 1
		}
		tent[j+1] = temp
	}
}

func urutTenantNama(tent *tabTenant, k *int) {
	var i, j, min int
	var temp tenant
	for i = 0; i < *k-1; i++ {
		min = i
		for j = i + 1; j < *k; j++ {
			if tent[j].nama < tent[min].nama {
				min = j
			}
		}
		temp = tent[i]
		tent[i] = tent[min]
		tent[min] = temp
	}
}

func printTenantData(tent tabTenant, j int) {
	var i int
	fmt.Println("Berikut tenant-tenant yang terdaftar :")
	i = 0
	for i < j {
		fmt.Printf("%d. %s\n", i+1, tent[i].nama)
		fmt.Println("jumlah transaksi :", tent[i].jumlahTransaksi)
		fmt.Println("Jumlah uang yang diterima tenant :", tent[i].uangTenant)
		fmt.Println("Jumlah Uang yang diterima admin pertransaksi :", tent[i].uangAdmin)
		fmt.Println("Berikut data transaksi tenant", tent[i].nama)
		for j := 0; j < tent[i].jumlahTransaksi; j++ {
			fmt.Println(j+1, tent[i].uangPerTransaksi[j])
		}
		i++
	}
	fmt.Println()
}
func printTenantDataSatu(tent tenant) {
	fmt.Println("Nama tenant : ", tent.nama)
	fmt.Println("Jumlah Transaksi : ", tent.jumlahTransaksi)
	for i := 0; i < tent.jumlahTransaksi; i++ {
		fmt.Println(i+1, tent.uangPerTransaksi[i])
	}
}
