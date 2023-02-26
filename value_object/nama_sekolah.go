package value_object

import "fmt"

type NamaSekolahDasar struct {
	value string
}

func NewNamaSekolahDasar(nama string) (*NamaSekolahDasar, error) {
	sd := NamaSekolahDasar{value: nama}
	err := sd.validate()
	if err != nil {
		return &NamaSekolahDasar{}, err
	}
	return &sd, nil
}

func (sd *NamaSekolahDasar) validate() error {
	// ada batasan jumlah karakter
	// depannya harus ada huruf sd
	strlen := len([]rune(sd.value))
	if strlen < 5 || strlen > 20 {
		return fmt.Errorf("Panjang karakter tidak sesuai ")
	}
	return nil
}
