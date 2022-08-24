package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	comando := 0
	fmt.Scan(&comando)

	fmt.Println("O comando selecionado foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo os logs...")
	} else if comando == 3 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Comando desconhecido")
	}

}
