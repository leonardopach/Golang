package main

import "fmt"

func mapsgo() {

	m := map[string]int{"sp": 1000000, "cg": 900000}
	makee := make(map[string]int)

	makee["sp"] = 100000
	makee["mg"] = 20000
	fmt.Println(m)
	fmt.Println(makee)

	for chave, valor := range m {
		fmt.Println("Cidade", chave, "H:", valor)
	}
}
