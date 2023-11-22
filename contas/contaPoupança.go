package contas

import (
	"github.com/JuanVinicius/Gerenciador/clientes"
)

type ContaPoupança struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupança) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque efetuado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupança) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "O deposito foi adicionado a conta, saldo atual é de", c.saldo
	} else {
		return "O valor é menor do que zero, saldo atual é de", c.saldo
	}
}

func (c *ContaPoupança) Obtersaldo() float64 {
	return c.saldo
}
