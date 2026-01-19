/*
Package trivialcompact é um pacote responsável por exportar funcionalidades
de compressão e descompressão para arquivos de textos que contenha sequências de DNAs.
*/
package trivialcompact

import (
	"bufio"
	"fmt"

	"github.com/Dom-Garotom/estruturaDeDados/compression/pkg"
)

/*
Compress é uma função responsável por compactar sequnências de DNAs em arquivos binários.
*/
func Compress(file *bufio.Scanner) {
	createdFile, _, err := pkg.CreateFile("../data/compressed/dna-compact", pkg.BINARY)

	if err != nil {
		return
	}
	defer createdFile.Close()

	var binaryValue = map[rune]uint8{
		'A': 0b00,
		'C': 0b01,
		'G': 0b10,
		'T': 0b11,
	}

	var byteSequency byte
	var bitCount uint8
	var buffer []byte
	var validSymbols byte

	for file.Scan() {
		for _, value := range file.Text() {
			encoding, ok := binaryValue[value]

			if !ok {
				continue
			}

			byteSequency = (byteSequency << 2) | encoding
			bitCount += 2

			if bitCount == 8 {
				buffer = append(buffer, byteSequency)
				byteSequency = 0
				bitCount = 0
			}
		}
	}

	if bitCount > 0 {
		validSymbols = bitCount / 2
		byteSequency <<= 8 - bitCount
		buffer = append(buffer, byteSequency)
	}

	createdFile.Write([]byte{validSymbols}) // header
	createdFile.Write(buffer)               // dados compactados
}

/*
Decompress é uma função responsável por descompactar arquivos
binários e transforma-los em arquivos de textos que contém uma sequnências de DNA.
*/
func Decompress(file []byte) {
	if len(file) < 2 {
		fmt.Printf("\nArquivo sem os headers necessários para a descompactação\n")
		return
	}

	osFile, writer, err := pkg.CreateFile("../data/decompressed/dna-decompressed", pkg.TXT)

	if err != nil || file == nil || writer == nil {
		return
	}

	var binaryValue = map[byte]string{
		0b00: "A",
		0b01: "C",
		0b10: "G",
		0b11: "T",
	}

	validSymbols := file[0]
	payload := file[1:]

	for index, currentByte := range payload {
		limit := 4

		if index == len(payload)-1 {
			limit = int(validSymbols)
		}

		for j := 0; j < limit; j++ {
			shift := 6 - (j * 2)
			bits := (currentByte >> shift) & 0b11
			writer.WriteString(binaryValue[bits])
		}
	}

	writer.Flush()
	defer osFile.Close()
}
