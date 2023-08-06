package model

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
}

func MostrarStruct() {

	endereco := Endereco{
		Rua:    "Rua X",
		Numero: 15,
		Cidade: "Campo Grande",
	}

	fmt.Println(endereco)
	endereco.Numero = 19
	fmt.Println(endereco)
}
