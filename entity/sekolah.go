package entity

import "latihan-unit-test-golang/value_object"

type SekolahDasar struct {
	namaSekolah value_object.NamaSekolahDasar
	status      string
	kelas       []Kelas
	alamat      string
	akreditasi  string
}

type Kelas struct {
	name     string
	properti Properti
	murid    []Murid
}

type Properti struct {
	nama   string
	jumlah int
}

type Murid struct {
	nama         string
	nomorAbsen   int
	jenisKelamin string
}

func NewSekolah(nama string, status string, alamat string, akreditasi string) {

}
