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
		fmt.Println("Uso: removehttp <file.txt>")
		return
	}

	// Nome do arquivo a ser processado
	fileName := os.Args[1]

	// Abre o arquivo para leitura
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Armazena as linhas processadas sem os prefixos
	var linhasSemProtocolo []string

	// Lê o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := scanner.Text()

		// Remove http:// ou https:// da linha
		linhaSemProtocolo := strings.ReplaceAll(linha, "http://", "")
		linhaSemProtocolo = strings.ReplaceAll(linhaSemProtocolo, "https://", "")

		// Adiciona a linha processada à lista
		linhasSemProtocolo = append(linhasSemProtocolo, linhaSemProtocolo)
	}

	// Verifica se houve erros durante a leitura
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		return
	}

	// Reabre o arquivo para escrita e sobrescreve o conteúdo
	file, err = os.Create(fileName)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo para escrita: %v\n", err)
		return
	}
	defer file.Close()

	// Escreve as linhas processadas de volta no arquivo
	writer := bufio.NewWriter(file)
	for _, linha := range linhasSemProtocolo {
		_, err := writer.WriteString(linha + "\n")
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
			return
		}
	}

	// Garante que todos os dados foram escritos no arquivo
	writer.Flush()

	fmt.Println("Arquivo atualizado com sucesso!")
}
