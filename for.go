package main

import "fmt"

func forGo() {

	texto := "palavra"
	for i := 0; i < len(texto); i++ {
		fmt.Println(string(texto[i]))
	}
}
