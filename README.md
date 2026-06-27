# 🌐 CLI Web Monitor (Desafio Go)

Um sistema de linha de comando (CLI) simples e eficiente desenvolvido em Go para monitorar a disponibilidade de websites. O programa lê uma lista de URLs de um arquivo texto, realiza requisições HTTP para validar o status code de cada site e armazena o histórico de varreduras em um arquivo de log local.

---

## 🛠️ Funcionalidades

* **Menu Interativo:** Fluxo contínuo de execução que permite ao usuário escolher entre iniciar o monitoramento, exibir os logs de histórico ou sair do programa com segurança.
* **Leitura Dinâmica de Sites:** Carregamento de URLs através de um arquivo de configuração externo (`sites.txt`).
* **Automação de Varreduras:** Execução de ciclos configuráveis de monitoramento com intervalos de tempo definidos (*delay*).
* **Persistência de Logs:** Histórico de status gravado em tempo real no arquivo `log.txt`, contendo data, hora, URL e o resultado do teste.
* **Tratamento de Erros Resiliente:** Tratamento nativo para cenários onde arquivos não existem ou conexões de rede falham, evitando que a aplicação pare de funcionar de forma inesperada.

---

## 🔬 Conceitos de Go Aplicados

Este projeto foi desenvolvido aplicando fundamentos e boas práticas do ecossistema Go:
* **Tratamento Explícito de Erros:** Ausência de blocos *try/catch*, utilizando o padrão idiomático do Go de validação de erros (`if err != nil`).
* **Gerenciamento de Arquivos Avançado:** Uso de máscaras de bits (`os.O_RDWR | os.O_CREATE | os.O_APPEND`) para manipulação segura de arquivos.
* **Leitura Otimizada com Buffer:** Utilização do pacote `bufio` para leitura eficiente de arquivos linha por linha sem sobrecarga de memória.
* **Conversão e Formatação Explicita de Tipos:** Uso de `strconv` e manipulação de slices de bytes (`[]byte`).
