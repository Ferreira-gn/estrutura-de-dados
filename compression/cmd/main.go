package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Dom-Garotom/estruturaDeDados/compression/pkg"
	trivialcompact "github.com/Dom-Garotom/estruturaDeDados/compression/trivial_compact"
)

func main() {

	file, err := os.Open("./data/dna.txt")

	if err != nil {
		fmt.Printf("\nErro ao abri o arquivo de dna\n")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	trivialcompact.Compress(scanner)

	compressedFile, err := os.ReadFile("./data/compressed/dna-compact.bin")

	if err != nil {
		fmt.Printf("\nErro ao abri o arquivo de dna\n")
		return
	}

	trivialcompact.Decompress(compressedFile)

	pkg.CompressionStatus("./data/dna.txt", "./data/compressed/dna-compact.bin")
	pkg.DiffChecker("./data/dna.txt", "./data/decompressed/dna-decompressed.txt")
}
