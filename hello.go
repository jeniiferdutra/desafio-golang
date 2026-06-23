package main

import "fmt"
import "os"

func main() {

	exibeIntroducao()
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
}