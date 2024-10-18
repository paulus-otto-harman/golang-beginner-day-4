package nasabah

import (
	"errors"
	"fmt"
)

type Akun struct {
	nama  string
	email string
	saldo Saldo
}

var daftarNasabah []Akun

func DaftarNasabah() []Akun {
	return daftarNasabah
}

func TambahAkun(nama string, email string) string {
	akun, err := Akun{}.InitAkun(nama, email)
	if err != nil {
		return pesan("", "akun gagal ditambahkan", err)
	}
	daftarNasabah = append(daftarNasabah, akun)
	return pesan("akun berhasil ditambahkan", "", nil)
}

func CariYangBernama(nama string) (*Akun, error) {
	for index, customer := range daftarNasabah {
		if customer.GetNama() == nama {
			return &daftarNasabah[index], nil
		}
	}
	return &Akun{}, errors.New("Nasabah Tidak Ditemukan")
}

func (akun Akun) InitAkun(nama string, email string) (Akun, error) {
	if nama == "" || email == "" {
		return Akun{}, errors.New("Nama dan Email wajib diisi")
	}
	akun.nama = nama
	akun.email = email
	return akun, nil
}

func (akun Akun) GetNama() string {
	return akun.nama
}

func (akun *Akun) SetorDana(uang int) string {
	_, err := akun.saldo.Kredit(uang)
	return pesan("saldo berhasil ditambahkan", "saldo tetap atau tidak bertambah", err)
}

func (akun *Akun) TarikDana(uang int) string {
	_, err := akun.saldo.Debet(uang)
	return pesan("saldo berhasil dikurangi", "saldo gagal dikurangi", err)
}

func pesan(pesanSukses string, pesanError string, err error) string {
	const Bold = "\033[1m%s\033[0m"
	const Color = "\x1b[%dm%s\x1b[0m"
	const Red = 31
	const Green = 92

	if err != nil {
		fmt.Printf(Color, Red, fmt.Sprintf("%s%s%s\n", "*** Error: ", err, " ***"))
		return fmt.Sprintf(Bold, pesanError)
	}
	return fmt.Sprintf(Color, Green, pesanSukses)
}
