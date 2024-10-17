package main

import (
	"20241017/pre/1/kendaraan"
	"fmt"
)

func main() {
	saya := Manusia{nama: "Amir"}
	fmt.Println(saya)

	dia := Person{Name: "Budi", age: 10}
	fmt.Println(dia)
	dia.Greet()
	dia.setAge(11)
	fmt.Println(dia)

	mobil := kendaraan.Mobil{}
	fmt.Println(mobil)
}
