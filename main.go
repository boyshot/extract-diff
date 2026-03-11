package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Verificar se foram passados 2 arquivos como parâmetros
	if len(os.Args) < 3 {
		fmt.Println("Uso: programa <arquivo1> <arquivo2>")
		os.Exit(1)
	}

	var err = createPaths("arqs")
	if err != nil {
		fmt.Printf("não foi possível criar args")
		os.Exit(1)
	}

	err = createPaths("results")
	if err != nil {
		fmt.Printf("não foi possível criar results")
		os.Exit(1)
	}

	arquivo1 := "./arqs/" + os.Args[1]
	arquivo2 := "./arqs/" + os.Args[2]

	// Ler o segundo arquivo e armazenar em um mapa para busca rápida
	linhasArquivo2, err := lerArquivo(arquivo2)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo2: %v\n", err)
		os.Exit(1)
	}

	// Converter para mapa para buscas O(1)
	mapa := make(map[string]bool)
	for _, linha := range linhasArquivo2 {
		mapa[linha] = true
	}

	// Ler o primeiro arquivo e encontrar linhas que não existem no segundo
	linhasArquivo1, err := lerArquivo(arquivo1)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo1: %v\n", err)
		os.Exit(1)
	}

	// Encontrar linhas que não existem
	var linhasDiff []string
	for _, linha := range linhasArquivo1 {
		if !mapa[linha] {
			linhasDiff = append(linhasDiff, linha)
		}
	}

	// Criar nome do arquivo com data/hora
	nomeArqDiff := "./results/diff__" + time.Now().Format("2006-01-02_15-04")

	// Escrever arquivo diff
	err = escreverArquivo(nomeArqDiff, linhasDiff)
	if err != nil {
		fmt.Printf("Erro ao escrever arquivo diff: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Arquivo criado com sucesso: %s\n", nomeArqDiff)
	fmt.Printf("Linhas diferentes encontradas: %d\n", len(linhasDiff))
}

// lerArquivo lê um arquivo e retorna um slice com suas linhas
func lerArquivo(caminho string) ([]string, error) {
	arquivo, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return linhas, nil
}

// escreverArquivo escreve um slice de linhas em um arquivo
func escreverArquivo(caminho string, linhas []string) error {

	if _, err := os.Stat(caminho); os.IsNotExist(err) {
		err = os.MkdirAll("./results", os.ModePerm)
		if err != nil {
			return err
		}
	}

	arquivo, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	writer := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		_, err := writer.WriteString(linha + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func createPaths(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err = os.MkdirAll("./"+name, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
