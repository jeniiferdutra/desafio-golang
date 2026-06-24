package main

import "fmt"
import "os"
import "net/http"

func main() {

	exibeIntroducao()

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
	sites := []string{"https://www.skoob.com.br/pt", "https://www.martinsfontespaulista.com.br/", "https://br.pinterest.com/", "https://studioghibli.com.br/studioghibli/"}
	
	for i, site := range sites{
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:" , resp.StatusCode)
	}
}