package main

import (
	"bufio"
	"fmt" // exibição de textos
	"io"
	"net/http" // requisições e servidores web
	"os"       // interação com o sistema operacional
	"strings"
	"time" // manipulação de datas e horários
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	registraLog("teste.com.br", false)

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

	fmt.Println("Olá, ", nome, "!")
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
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:" , resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites[]string // slice vazio

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err) 
	}

	leitor := bufio.NewReader(arquivo) // vai retornar um leitor

	for {
		linha, err := leitor.ReadString('\n') // ler até a quebra de linha 
		linha = strings.TrimSpace(linha)
	
		sites = append(sites, linha) // vou add no slice vazio

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)
	arquivo.Close()
}