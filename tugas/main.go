package main

import (
	"20241017/tugas/nasabah"
	"fmt"
)

func main() {
	fmt.Printf("%s %+v\n", nasabah.TambahAkun("satu", "satu@mail.com"), nasabah.DaftarNasabah())
	fmt.Printf("%s %+v\n", nasabah.TambahAkun("dua", "dua@mail.com"), nasabah.DaftarNasabah())
	fmt.Printf("%s %+v\n", nasabah.TambahAkun("tiga", ""), nasabah.DaftarNasabah())

	customer, err := nasabah.CariYangBernama("satu")
	if err != nil {
		fmt.Printf("Error Pencarian : %s\n", err)
		return
	}

	fmt.Printf("%s %+v\n", customer.SetorDana(0), nasabah.DaftarNasabah())
	fmt.Printf("%s %+v\n", customer.SetorDana(10), nasabah.DaftarNasabah())
	fmt.Printf("%s %+v\n", customer.TarikDana(15), nasabah.DaftarNasabah())
	fmt.Printf("%s %+v\n", customer.TarikDana(3), nasabah.DaftarNasabah())

}
