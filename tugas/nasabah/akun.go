package nasabah

import "fmt"

type Akun struct {
	nama  string
	email string
	saldo Saldo
}

func (akun Akun) InitAkun(nama string, email string) Akun {
	akun.nama = nama
	akun.email = email
	return akun
}

func (akun Akun) GetNama() string {
	return akun.nama
}

func (akun *Akun) SetorDana(uang int) string {
	_, err := akun.saldo.Kredit(uang)
	if err != nil {
		fmt.Printf("\x1b[%dm%s%s%s\x1b[0m\n", 31, "*** Error: ", err, " ***")
		return "saldo tetap atau tidak bertambah"

	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, "saldo berhasil ditambahkan")
}

func (akun *Akun) TarikDana(uang int) string {
	_, err := akun.saldo.Debet(uang)
	if err != nil {
		fmt.Printf("\x1b[%dm%s%s%s\x1b[0m\n", 31, "*** Error: ", err, " ***")
		return fmt.Sprintf("saldo gagal dikurangi")

	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, "saldo berhasil dikurangi")
}
