package main

import (
	contas "banco/contas"
	"fmt"
)

type iConta interface {
	Sacar(valordoSaque float64) (string, float64)
}

func PagarBoleto(conta iConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {
	var contadoGuilerme contas.ContaCorrente = contas.ContaCorrente{}
	contadoGuilerme.Titular.Nome = "Guilherme"
	contadoGuilerme.Titular.CPF = "11122233301"
	contadoGuilerme.Titular.Profissao = "Engenheiro de Software"
	contadoGuilerme.NumeroAgencia = "123-1"
	contadoGuilerme.NumeroConta = "112233-01"
	contadoGuilerme.Depositar(1125.50)

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular.Nome = "Cris"
	contaDaCris.Titular.CPF = "11122233302"
	contaDaCris.Titular.Profissao = "Product Owner"
	contaDaCris.NumeroAgencia = "123-1"
	contaDaCris.NumeroConta = "112233-02"
	contaDaCris.Depositar(1500.00)

	// fmt.Println(contaDaCris.sacar(150))
	// fmt.Println(contaDaCris.depositar(1250.50))

	fmt.Println("Saldo conta da Cris ", contaDaCris.ObterSaldo())
	fmt.Println("Saldo conta do Guilherme ", contadoGuilerme.ObterSaldo())

	if contaDaCris.Transferir(500, &contadoGuilerme) {
		fmt.Println("Saldo conta da Cris ", contaDaCris.ObterSaldo())
		fmt.Println("Saldo conta do Guilherme ", contadoGuilerme.ObterSaldo())
	} else {
		fmt.Println("Ocorreu um problema. Transferência não realizada")
	}

	PagarBoleto(contaDaCris, 100)
	fmt.Println("Saldo conta da Cris ", contaDaCris.ObterSaldo())

	// fmt.Println("Soma", contaDaCris.soma(500, 300, 100, 100))

	// fmt.Println(contadoGuilerme, "type: ", reflect.TypeOf(contadoGuilerme))
	// fmt.Println(contaDaBruna, "type: ", reflect.TypeOf(contaDaBruna))
	// fmt.Println(*contaDaCris, "type: ", reflect.TypeOf(contaDaCris))

}
