package pkg

import (
	"fmt"
	"os"
)

/*
CompressionStatus é uma função responsável por retornar os status da compressão de um arquivo,
comparando o arquivo original e sua versão comprimida, informando os seus tamanhos
e se houve um ganho significativo na economia do armazenamento.
*/
func CompressionStatus(originalFileName string, compressedFileName string) {
	originalFile, errOnOpenOriginalFile := os.Stat(originalFileName)
	compressedFile, errOnOpenCompressedFile := os.Stat(compressedFileName)

	if errOnOpenCompressedFile != nil || errOnOpenOriginalFile != nil {
		fmt.Printf("Erro ao abrir arquivos de textos para a verificação de status da compressão")
		return
	}

	originalFileSize := FileSizeFormatter(originalFile.Size())
	compressedFileSize := FileSizeFormatter(compressedFile.Size())

	fmt.Printf("\n\n\033[1;33mVerificação dos status da manutenção\033[0m\n\n")

	fmt.Printf("Nome %31s Tamanho\n\n", "")
	fmt.Printf("%s %27s\n", "Arquivo original", originalFileSize)
	fmt.Printf("%s %25s\n", "Arquivo compactado", compressedFileSize)

	economy := (compressedFile.Size() / originalFile.Size()) * 100

	fmt.Printf("\033[1m\n\nResultado :\n\033[0m")

	if economy <= 0 {
		fmt.Printf("\033[31mNenhum ganho significativo na econômia de espaço.\033[0m\n\n")
		return
	}

	fmt.Printf("\n\033[32mEconomia de %d%s", economy, "% porcentos\033[0m\n\n")
}
