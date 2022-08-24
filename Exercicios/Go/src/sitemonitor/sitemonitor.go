package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var comando int
	for {
		exibeMenu()
		lerComando(&comando)
		executaComando(&comando)
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando(comandoLido *int) {
	fmt.Scan(comandoLido)
}

func executaComando(comandoLido *int) {
	var url string

	fmt.Println("O comando selecionado foi", *comandoLido)

	switch *comandoLido {
	case 1:
		fmt.Println("Qual site você quer monitorar?")
		fmt.Scan(&url)
		fmt.Println("Monitorando", url, "...")
		statCode := monitoraSite(url)
		if statCode == 200 {
			fmt.Println("Site", url, "carregado com sucesso!")
		} else {
			fmt.Println("Site", url, "com problema, status:", statCode)
		}
	case 2:
		fmt.Println("Exibindo os logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}
}

func monitoraSite(url string) int {
	resp, _ := http.Get(url)
	return resp.StatusCode
}
