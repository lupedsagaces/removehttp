package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Verifica se o arquivo foi passado como argumento
	if len(os.Args) < 2 {
		fmt.Println("Usage: removehttp <file.txt>")
		return
	}

	// Abre o arquivo
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	// Processa cada linha
	for scanner.Scan() {
		linha := scanner.Text()

		// Remove http:// ou https:// da linha
		linhaSemProtocolo := strings.ReplaceAll(linha, "http://", "")
		linhaSemProtocolo = strings.ReplaceAll(linhaSemProtocolo, "https://", "")

		// Exibe o resultado
		fmt.Println(linhaSemProtocolo)
	}

	// Verifica se houve erros durante a leitura do arquivo
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
	}
}
