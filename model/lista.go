package model

import "fmt"

func listas() {
	lista := []int{4, 9, 6, 7}

	fmt.Println("Lista: ", lista)
	fmt.Println("Lista1: ", lista[0])

	lista = append(lista, 8, 10)

	for i := 0; i < len(lista); i++ {
		fmt.Println(lista[i])
	}
	listaString := make([]string, 0)
	listaString = append(listaString, "hello", "moto")

	fmt.Println(listaString)
}
