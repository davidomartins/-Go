package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "David"
	var versao = 1.1
	fmt.Println("Olá mundo!")
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão desse programa é", versao)

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))

}
