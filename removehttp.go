package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var scanner *bufio.Scanner

	// Verifica se o arquivo foi passado como argumento
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		// Caso contrário, lê da entrada padrão (pipe)
		scanner = bufio.NewScanner(os.Stdin)
	}

	// Processa cada linha
	for scanner.Scan() {
		linha := scanner.Text()

		// Remove http:// ou https:// da linha
		linhaSemProtocolo := strings.ReplaceAll(linha, "http://", "")
		linhaSemProtocolo = strings.ReplaceAll(linhaSemProtocolo, "https://", "")

		// Exibe o resultado
		fmt.Println(linhaSemProtocolo)
	}

	// Verifica se houve erros durante a leitura
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler a entrada: %v\n", err)
	}
}
