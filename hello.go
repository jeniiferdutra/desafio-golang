package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	leSitesDoArquivo()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1: 
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")	
			os.Exit(0) // Saiu com sucesso
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) // Ocorreu algum erro no código
		}
	}

}

func exibeIntroducao() {
	nome := "Jenifer"
	versao := 1.26

	fmt.Println("Olá, ", nome)
	fmt.Println("Versão atual do programa", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)	

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.skoob.com.br/pt", "https://www.martinsfontespaulista.com.br/", "https://br.pinterest.com/", "https://studioghibli.com.br/studioghibli/"}
	
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos ; i++ {
		for i, site := range sites{
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}	
		time.Sleep(delay * time.Second)
		fmt.Println("")
	} 

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:" , resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites[]string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err) 
	}

	fmt.Println(arquivo)
	return sites
}