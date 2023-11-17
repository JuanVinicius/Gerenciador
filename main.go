package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaJuan := contaCorrente{
		"Juan",
		222,
		111223,
		120.50,
	}

	contaRebecca := contaCorrente{
		"Rebecca",
		222,
		111222,
		120.50,
	}

	fmt.Println(contaJuan, contaRebecca)
	contaJuan.Transferir(100.0, &contaRebecca)
	fmt.Println(contaJuan, contaRebecca)
}

func (c *contaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque efetuado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *contaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "O deposito foi adicionado a conta, saldo atual é de", c.saldo
	} else {
		return "O valor é menor do que zero, Saldo atual é de", c.saldo
	}
}

func (c *contaCorrente) Transferir(valorDaTransferencia float64, contaDeDestino *contaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDeDestino.Depositar(valorDaTransferencia)
		fmt.Println("Transferência realizada com sucesso")
		return true
	} else {
		fmt.Println("Saldo insuficiente")
		return false
	}
}
