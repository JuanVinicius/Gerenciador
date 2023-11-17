package contas

import (
	"fmt"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.Saldo {
		c.Saldo -= valorDoSaque
		return "Saque efetuado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "O deposito foi adicionado a conta, Saldo atual é de", c.Saldo
	} else {
		return "O valor é menor do que zero, Saldo atual é de", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDeDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		fmt.Println("Transferência realizada com sucesso")
		return true
	} else {
		fmt.Println("Saldo insuficiente")
		return false
	}
}
