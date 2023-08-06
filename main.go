package main

import (
	"fmt"
	"golanestudo/model"
)

func main() {
	fmt.Println("Iniciando...")

	automovelMoto := model.Automovel{
		Ano:    2023,
		Placa:  "Honda",
		Modelo: "Cg",
	}

	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 125,
	}

	fmt.Println(moto)
	fmt.Println(moto.Ano)
}
