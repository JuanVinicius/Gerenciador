package main

import (
	"fmt"

	"github.com/JuanVinicius/Gerenciador/contas"
)

func main() {

	contaJuan := contas.ContaCorrente{
		Titular:       "Juan",
		NumeroAgencia: 222,
		NumeroConta:   111223,
		Saldo:         120.50,
	}

	contaRebecca := contas.ContaCorrente{
		Titular:       "Rebecca",
		NumeroAgencia: 222,
		NumeroConta:   111222,
		Saldo:         120.5,
	}

	fmt.Println(contaJuan, contaRebecca)
	contaJuan.Transferir(100.0, &contaRebecca)
	fmt.Println(contaJuan, contaRebecca)
}
