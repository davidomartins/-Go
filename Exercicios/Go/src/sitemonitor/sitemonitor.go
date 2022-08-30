package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}
}

func statusCodeDoSite(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		return 500
	} else {
		return resp.StatusCode
	}
}

func listaDeSites() []string {
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br/", "https://www.caelum.com.br/xpto", "https://www.xpto.x.pto/xpto"}
	return sites
}

func monitorarSites() {

	// listaDeSite := listaDeSites()
	listaDeSite := lerSitedoArquivo("sites.txt")

	fmt.Println("Monitorando:")

	for _, url := range listaDeSite {
		statCode := statusCodeDoSite(url)
		if statCode == 200 {
			fmt.Println("Site", url, "carregado com sucesso!")
			registraLog(url, true)
		} else {
			fmt.Println("Site", url, "com problema, status:", statCode)
			registraLog(url, false)
		}
	}
}

func lerSitedoArquivo(caminho_arquivo string) []string {

	arq, err1 := os.Open(caminho_arquivo)

	if err1 != nil {
		fmt.Println("Não foi possível abrir o arquivo ", caminho_arquivo, ". Detalhe do erro: ", err1)
		os.Exit(-1)
	}

	leitor := bufio.NewReader(arq)

	linha, err2 := leitor.ReadString('\n')

	if err2 != nil {
		if err2 == io.EOF {
			fmt.Println("O arquivo", caminho_arquivo, "está vazio!")
		} else {
			fmt.Println("Ocorreu um erro ao ler o arquivo. Detalhe do erro: ", err2)
		}
		arq.Close()
		os.Exit(-1)
	}

	var sites []string

	for {
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err2 == io.EOF {
			break
		}
		linha, err2 = leitor.ReadString('\n')
	}
	arq.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
