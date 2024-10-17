package main

import (
	"20241017/class/1/nasabah"
	"errors"
	"fmt"
)

var daftarNasabah []nasabah.Akun

func TambahNasabah(nama string, email string) {
	akun := nasabah.Akun{}.InitAkun(nama, email)
	daftarNasabah = append(daftarNasabah, akun)
	fmt.Printf("akun berhasil ditambahkan %+v\n", daftarNasabah)
}

func CariNasabah(nama string) (*nasabah.Akun, error) {
	for index, customer := range daftarNasabah {
		if customer.GetNama() == nama {
			return &daftarNasabah[index], nil
		}
	}
	return &nasabah.Akun{}, errors.New("Tidak Ditemukan")
}

func main() {
	TambahNasabah("satu", "satu@mail.com")
	TambahNasabah("dua", "dua@mail.com")
	//tambahNasabah("tiga", "tiga@mail.com")

	customer, err := CariNasabah("dua")
	if err != nil {
		return
	}

	fmt.Printf("%s %+v\n", customer.SetorDana(0), daftarNasabah)
	fmt.Printf("%s %+v\n", customer.SetorDana(10), daftarNasabah)
	fmt.Printf("%s %+v\n", customer.TarikDana(15), daftarNasabah)
	fmt.Printf("%s %+v\n", customer.TarikDana(3), daftarNasabah)

}
