package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia string
	NumeroConta   string
	Saldo         float64
}

func (c *ContaCorrente) sacar(valordoSaque float64) (string, float64) {
	if valordoSaque < 0 {
		return "Ops! Esse valor é inválido!", valordoSaque
	} else {

		podeSacar := valordoSaque <= c.Saldo

		if podeSacar {
			c.Saldo -= valordoSaque
			return "Saque realizado com sucesso", c.Saldo
		} else {
			return "Saldo insuficiente", c.Saldo

		}
	}
}

func (c *ContaCorrente) depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito < 0 {
		return "Ops! Essa valor de depósito é inválido!", valordoDeposito
	} else {
		c.Saldo += valordoDeposito
		return "Depósito realizado com sucesso!", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valordaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valordaTransferencia < 0 {
		return false
	} else {

		podeSacar := valordaTransferencia <= c.Saldo

		if podeSacar {
			c.Saldo -= valordaTransferencia
			contaDestino.depositar(valordaTransferencia)
			return true
		} else {
			return false

		}
	}
}

// func (c *ContaCorrente) soma(numeros ...int64) int64 {
// 	var soma int64
// 	for _, numero := range numeros {
// 		soma += numero
// 	}
// 	return soma
// }
