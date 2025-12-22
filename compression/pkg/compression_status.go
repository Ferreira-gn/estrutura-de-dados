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
	originalFile, errOriginal := os.Stat(originalFileName)
	compressedFile, errCompressed := os.Stat(compressedFileName)

	if errOriginal != nil || errCompressed != nil {
		fmt.Println("Erro ao abrir arquivos para verificação de status da compressão.")
		return
	}

	originalSize := originalFile.Size()
	compressedSize := compressedFile.Size()

	if originalSize == 0 {
		fmt.Println("Arquivo original vazio. Não é possível calcular economia.")
		return
	}

	originalFormatted := FileSizeFormatter(originalSize)
	compressedFormatted := FileSizeFormatter(compressedSize)

	fmt.Printf("\n\n\033[1;33mVerificação dos status da compressão\033[0m\n\n")
	fmt.Printf("Nome %31s Tamanho\n\n", "")
	fmt.Printf("%-30s %10s\n", "Arquivo original", originalFormatted)
	fmt.Printf("%-30s %10s\n", "Arquivo compactado", compressedFormatted)

	ratio := float64(compressedSize) / float64(originalSize)
	economy := (1 - ratio) * 100

	fmt.Printf("\033[1m\n\nResultado:\n\033[0m")

	switch {
	case economy > 0:
		fmt.Printf("\033[32mEconomia de %.2f%% no armazenamento.\033[0m\n\n", economy)
	case economy == 0:
		fmt.Printf("\033[33mNenhuma economia de espaço obtida.\033[0m\n\n")
	default:
		fmt.Printf(
			"\033[31mArquivo compactado é %.2f%% maior que o original.\033[0m\n\n",
			-economy,
		)
	}
}
