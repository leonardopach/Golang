package model

import (
	"fmt"
	"reflect"
	"strconv"
)

func conversao() {
	var numero int16 = 127
	var numero2 float32
	texto := "42"
	numero2 = float32(numero)

	b, _ := strconv.ParseInt(texto, 10, 64)

	println(b)
	fmt.Println(reflect.TypeOf(numero2))
	fmt.Println(numero2)
}
