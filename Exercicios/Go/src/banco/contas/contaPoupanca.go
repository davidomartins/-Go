package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia string
	NumeroConta   string
	saldo         float64
	dataBase      string
}

func (c *ContaPoupanca) sacar(valordoSaque float64) (string, float64) {
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

func (c *ContaPoupanca) Depositar(valordoDeposito float64) (string, float64) {
	if valordoDeposito < 0 {
		return "Ops! Essa valor de depósito é inválido!", valordoDeposito
	} else {
		c.saldo += valordoDeposito
		return "Depósito realizado com sucesso!", c.saldo
	}

}

func (c *ContaPoupanca) Transferir(valordaTransferencia float64, contaDestino *ContaPoupanca) bool {
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

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
