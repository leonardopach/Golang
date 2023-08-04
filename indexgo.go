package main

import "fmt"

func indexgo() {
	var listatoda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	segundaLista := listatoda[0 : len(listatoda)-1]

	listatoda2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	listatoda2 = append(listatoda2, 10)
	fmt.Println(segundaLista)
	fmt.Println(listatoda2)
}
