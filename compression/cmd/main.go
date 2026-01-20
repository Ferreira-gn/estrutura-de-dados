package main

import (
	"bufio"
	"fmt"
	"os"

	performativecompression "github.com/Dom-Garotom/estruturaDeDados/compression/internal/performative_compression"
	trivialcompact "github.com/Dom-Garotom/estruturaDeDados/compression/internal/trivial_compact"
	"github.com/Dom-Garotom/estruturaDeDados/compression/pkg"
)

func main() {
	file, err := os.Open("../data/dna.txt")

	if err != nil {
		fmt.Printf("\nErro ao abri o arquivo de dna\n")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	trivialcompact.Compress(scanner)

	compressedFile, err := os.ReadFile("../data/compressed/dna-compact.bin")

	if err != nil {
		fmt.Printf("\nErro ao abri o arquivo compactado do dna\n")
		return
	}

	trivialcompact.Decompress(compressedFile)

	fmt.Printf("\n\n\033[1;34mBenchMark do modelo de compressão trivial    \n\n\033[0m ")
	pkg.CompressionStatus("../data/dna.txt", "../data/compressed/dna-compact.bin")
	pkg.DiffChecker("../data/dna.txt", "../data/decompressed/dna-decompressed.txt")

	performativecompression.Compress("../data/dna.txt", "../data/compressed/dna-performatic-compact")
	performativecompression.Decompress("../data/compressed/dna-performatic-compact.bin")

	fmt.Printf("\n\n\033[1;34mBenchMark do modelo de compressão performatico    \n\n\033[0m ")
	pkg.CompressionStatus("../data/dna.txt", "../data/compressed/dna-performatic-compact.bin")
	pkg.DiffChecker("../data/dna.txt", "../data/decompressed/dna-performatic-decompressed.txt")
}
