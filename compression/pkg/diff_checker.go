package pkg

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*
DiffChecker é uam função  responsável por verificar as diferenças e o nível de compatibilidade entre dois arquivos.
*/
func DiffChecker(originalFileName string, variantFileName string) {
	originalFile, errorOnReadOriginalFile := os.ReadFile(originalFileName)
	variantFile, errorOnReadVariantFile := os.ReadFile(variantFileName)

	var originalFileBuffer bytes.Buffer
	var variantFileBuffer bytes.Buffer

	if errorOnReadOriginalFile != nil || errorOnReadVariantFile != nil {
		fmt.Printf("\n\nErro ao tentar abrir os arquivos usados para verificação de igualdade\n\n")
		return
	}

	originalFileBuffer.Write(originalFile)
	variantFileBuffer.Write(variantFile)

	originalString := strings.ReplaceAll(originalFileBuffer.String(), "\n", "")
	variantString := strings.ReplaceAll(variantFileBuffer.String(), "\n", "")

	fmt.Printf("\033[1;33m\n\nVerificação de inconsistências entre os arquivos de texto\n\n\033[0m")

	fmt.Printf("Localização %-*s Referência\n\n", 34, "")
	fmt.Printf("%-*s | %s |\n", 45, originalFileName, "Original")
	fmt.Printf("%-*s | %s |\n", 45, variantFileName, "Variante")

	fmt.Printf("\033[1m\n\nResultado :\n\033[0m")

	if originalString == variantString {
		fmt.Printf("\033[32mOs arquivos de texto são identicos\n\n\033[0m")
		return
	}

	fmt.Printf("\033[31mOs arquivos de texto são diferentes\n\n\033[0m")
}
