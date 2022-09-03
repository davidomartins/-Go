package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia string
	NumeroConta   string
	saldo         float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) (string, float64) {
	if valordoSaque < 0 {
		return "Ops! Esse valor é inválido!", valordoSaque
	} else {

		podeSacar := valordoSaque <= c.saldo

		if podeSacar {
			c.saldo -= valordoSaque
			return "Saque realizado com sucesso", c.saldo
		} else {
			return "saldo insuficiente", c.saldo

		}
	}
}

func (c *ContaCorrente) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito < 0 {
		return "Ops! Essa valor de depósito é inválido!", valordoDeposito
	} else {
		c.saldo += valordoDeposito
		return "Depósito realizado com sucesso!", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valordaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valordaTransferencia < 0 {
		return false
	} else {

		podeSacar := valordaTransferencia <= c.saldo

		if podeSacar {
			c.saldo -= valordaTransferencia
			contaDestino.Depositar(valordaTransferencia)
			return true
		} else {
			return false

		}
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

// func (c *ContaCorrente) soma(numeros ...int64) int64 {
// 	var soma int64
// 	for _, numero := range numeros {
// 		soma += numero
// 	}
// 	return soma
// }
