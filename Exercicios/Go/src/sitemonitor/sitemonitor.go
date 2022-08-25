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
		executaComando(comando)
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

func executaComando(comandoLido int) {
	switch comandoLido {
	case 1:
		monitorarSites()
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

func statusCodeDoSite(url string) int {
	resp, _ := http.Get(url)
	return resp.StatusCode
}

func listaDeSites() []string {
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br/", "https://www.caelum.com.br/xpto"}
	return sites
}

func monitorarSites() {
	listaDeSite := listaDeSites()
	for i := 0; i < len(listaDeSite); i++ {
		url := listaDeSite[i]
		fmt.Println("Monitorando", url, "...")
		statCode := statusCodeDoSite(url)
		if statCode == 200 {
			fmt.Println("Site", url, "carregado com sucesso!")
		} else {
			fmt.Println("Site", url, "com problema, status:", statCode)
		}
	}
}
