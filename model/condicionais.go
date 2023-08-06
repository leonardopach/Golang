package model

import "fmt"

func condicionais() {
	salario := 999.00

	var salarioMaisOBonus float64

	salarioMaisOBonus = salario

	if salario < 1000 {
		salarioMaisOBonus = (salario + 100)
	} else {
		salarioMaisOBonus += 20
	}

	fmt.Println("Salario: ", salarioMaisOBonus)

}
