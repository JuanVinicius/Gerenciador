package main

import (
	"fmt"

	"github.com/JuanVinicius/Gerenciador/contas"
)

func main() {
	contaIgor := contas.ContaPoupança{}
	contaIgor.Depositar(100)
	PagarBoleto(&contaIgor, 60)
	fmt.Println(contaIgor.Obtersaldo())

	contaEster := contas.ContaCorrente{}
	contaEster.Depositar(100)
	PagarBoleto(&contaEster, 50)
	fmt.Println(contaEster.Obtersaldo())
}

func PagarBoleto(conta verificarConta, ValorDoBoleto float64) {
	conta.Sacar(ValorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
