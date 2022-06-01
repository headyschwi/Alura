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

const monitoramentos = 3
const delay = 5

func main() {
	introducao()
	for {
		exibeMenu()

		comando := capturaComando()
		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			exibeLog()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando.")
			os.Exit(-1)
		}
	}
}

func introducao() {
	nome := "Eder"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}
func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir log")
	fmt.Println("0- Sair do programa")
}
func capturaComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	fmt.Println("")
	return comando
}
func iniciaMonitoramento() {
	fmt.Println("Iniciando monitoramento...")

	sites := capturaSitesArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		fmt.Println("")
		time.Sleep(delay * time.Second)

	}
	fmt.Println("")
}
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Comunicação com", site, "concluída com sucesso!")
		escreveLog(site, true)
	} else {
		fmt.Println("Houve problemas na comunicação com", site, "StatusCode:", resp.StatusCode)
		escreveLog(site, false)
	}
}
func capturaSitesArquivo() []string {
	var sites []string
	site, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(site)

	for {
		linha, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}
	site.Close()
	return sites
}
func escreveLog(site string, online bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	date := time.Now().Format("02/01/2006 15:04:05")
	arquivo.WriteString(date + " - " + site + " - online: " + strconv.FormatBool(online) + "\n")
	arquivo.Close()
}
func exibeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}
