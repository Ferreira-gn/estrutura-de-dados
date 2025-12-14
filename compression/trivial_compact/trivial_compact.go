/*
Package trivialcompact é um pacote responsável por exportar funcionalidades
de compressão e descompressão para arquivos de textos que contenha sequências de DNAs.
*/
package trivialcompact

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type FileType string

const (
	TXT    FileType = ".txt"
	BINARY FileType = ".bin"
)

/*
Compress é uma função responsável por compactar sequnências de DNAs em arquivos binários.
*/
func Compress(file *bufio.Scanner) {
	createdFile, _, err := createFile("./data/compressed/dna-compact", BINARY)

	if err != nil || file == nil || createdFile == nil {
		return
	}

	var binaryValue = map[rune]byte{
		'A': 00,
		'C': 01,
		'G': 10,
		'T': 11,
	}

	var byteSequency byte


		// ler todo o arquivo de texto e trasnformamos os seus valores em bytes.
		// pegamos esse array de bytes e pegamos seus caracteres e transformamos em um um valor binário, isso pulando de de 4 em 4.
		// depois escrevemos os seus valores em binário.


	for file.Scan() {
		for _, value := range file.Text() {



			err := binary.Write(createdFile, binary.LittleEndian, binaryValue[value])

			if err != nil {
				fmt.Printf("Error ao salvar o valor %v no arquivo %v\n", value, err)
				continue
			}
		}
	}

	defer createdFile.Close()
}

/*
Decompress é uma função responsável por descompactar arquivos
binários e transforma-los em arquivos de textos que contém uma sequnências de DNA.
*/
func Decompress(file []byte) {
	osFile, writer, err := createFile("./data/decompressed/dna-decompressed", TXT)

	if err != nil || file == nil || writer == nil {
		return
	}

	var binaryValue = map[byte]string{
		00: "A",
		01: "C",
		10: "G",
		11: "T",
	}

	for _, byteValue := range file {
		_, err := writer.WriteString(binaryValue[byteValue])

		if err != nil {
			fmt.Printf("Erro ao salvar valor decodificado %v", err)
		}
	}

	writer.Flush()
	defer osFile.Close()
}

func createFile(fileName string, fileType FileType) (*os.File, *bufio.Writer, error) {
	file, err := os.Create(fileName + string(fileType))

	if err != nil {
		fmt.Printf("\nErro ao criar o arquivo\n")
		fmt.Printf("%v", err)
		return nil, nil, err
	}

	if fileType == TXT {
		writer := bufio.NewWriter(file)
		return file, writer, nil
	}

	if fileType == BINARY {
		return file, nil, nil
	}

	return nil, nil, err
}
