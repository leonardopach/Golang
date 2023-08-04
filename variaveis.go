package main

import (
	"fmt"
	"reflect"
)

func main() {
	var peso int
	peso = 10
	var mensagem string
	mensagem = "10 de peso"
	mensagem2 := "test"

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(peso)
	fmt.Println(reflect.TypeOf(peso))
}
