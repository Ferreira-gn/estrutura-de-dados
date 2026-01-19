package performativecompression

import (
	"fmt"

	"github.com/Dom-Garotom/estruturaDeDados/compression/pkg"
)

func Compress(path string) {
	dna, erroOnReadDNAFile := pkg.ReadTextFiles(path)

	if erroOnReadDNAFile != nil {
		fmt.Printf("Erro ao tentar abrir o arquivo de texto do DNA\nError : %v", erroOnReadDNAFile)
		return
	}

	formattedDNA, dnaSequences, _ := formatDNA(dna)
	buffer, errOnConvertDNAToBinary := convertDNAToBinary(formattedDNA, dnaSequences)

	if errOnConvertDNAToBinary != nil {
		fmt.Printf("Erro ao tentar converter o arquivo de texto do DNA para binário\nError : %v", errOnConvertDNAToBinary)
		return
	}

	// salva o buffer com os dados em binário em um arquivo binário
	createdFile, _, err := pkg.CreateFile("../data/compressed/dna-performatic-compact", pkg.BINARY)

	if err != nil {
		return
	}
	
	createdFile.Write(buffer)
	defer createdFile.Close()
}
