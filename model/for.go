package model

import "fmt"

func ForGo() {

	texto := "palavra"
	for i := 0; i < len(texto); i++ {
		fmt.Println(string(texto[i]))
	}
}
