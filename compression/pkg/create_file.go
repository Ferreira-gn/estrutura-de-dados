package pkg

import (
	"bufio"
	"fmt"
	"os"
)

type FileType string

const (
	TXT    FileType = ".txt"
	BINARY FileType = ".bin"
)

// CreateFile é uma função responsável por criar um arquivo de texto ou binário.
//
// Ele retorna :
//
// - Um ponteiro para o arquivo criado.  *os.File
//
// - Um ponteiro para o escritor de buffer. *bufio.Writer
//
// - Um erro. error
func CreateFile(fileName string, fileType FileType) (*os.File, *bufio.Writer, error) {
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
