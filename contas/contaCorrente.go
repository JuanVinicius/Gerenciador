package contas

import (
	"fmt"

	"github.com/JuanVinicius/Gerenciador/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque efetuado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "O deposito foi adicionado a conta, saldo atual é de", c.saldo
	} else {
		return "O valor é menor do que zero, saldo atual é de", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDeDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		fmt.Println("Transferência realizada com sucesso")
		return true
	} else {
		fmt.Println("saldo insuficiente")
		return false
	}
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}
