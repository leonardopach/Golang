package model

import (
	"fmt"
	"reflect"
)

func variaveis() {
	var peso int
	peso = 10
	var mensagem string = "mensagem aqui"
	mensagem2 := "test"
	var result = 12

	const valor = 24

	fmt.Println(result)
	fmt.Println(valor)
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(peso)
	fmt.Println(reflect.TypeOf(peso))
}
