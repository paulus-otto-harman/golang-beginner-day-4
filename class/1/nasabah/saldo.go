package nasabah

import "errors"

type Saldo struct {
	saldo int
}

func (saldo *Saldo) Tambah(duit int) (int, error) {
	if duit <= 0 {
		return saldo.saldo, errors.New("jumlah setoran tidak boleh lebih kecil atau sama dengan 0")
	}

	saldo.saldo += duit
	return saldo.saldo, nil
}

func (saldo *Saldo) Debet(duit int) (int, error) {
	if saldo.saldo < duit {
		return saldo.saldo, errors.New("saldo tidak mencukupi")
	}
	saldo.saldo -= duit
	return saldo.saldo, nil
}
