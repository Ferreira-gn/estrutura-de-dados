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



func createFile(fileName string, fileType FileType) (*os.File, *bufio.Writer, error) {
	file, err := os.Create(fileName + string(fileType))

	if err != nil {
		fmt.Printf("\nErro ao criar o arquivo\n")
		fmt.Printf("%v" , err)
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
