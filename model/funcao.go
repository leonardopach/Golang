package model

import "fmt"

func init() {
	fmt.Println("Função init")
}

func funcao() {

	ImprimirMensagem("Hello moto")

	fmt.Println(Soma(10, 2))

	soma, subtracao, divisao, multiplicacao := Operacao(1, 2)
	fmt.Println(soma, subtracao, divisao, multiplicacao)
}

func ImprimirMensagem(mensagem string) {
	fmt.Println(mensagem)
}

func Soma(x int, y int) int {
	result := x + y
	return result
}

func Operacao(numero int, numero2 int) (somar int, subtracao int, divisao int,
	multiplicacao int) {
	somar = numero + numero2
	subtracao = numero - numero2
	divisao = numero / numero2
	multiplicacao = numero * numero2
	return
}
