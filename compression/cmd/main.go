package main

import (
	"bufio"
	"fmt"
	"os"

	trivialcompact "github.com/Dom-Garotom/estruturaDeDados/compression/trivial_compact"
)

func main() {

	file, err := os.Open("data/dna.txt")

	if err != nil {
		fmt.Printf("\nErro ao abri o arquivo de dna\n")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	trivialcompact.Compress(scanner)
}
