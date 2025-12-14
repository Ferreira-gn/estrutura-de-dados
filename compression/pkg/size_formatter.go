/*
Package pkg é um pacote que fornece diversas funções 
utilitárias que podem ser usadas por qualquer pacote externo.
*/
package pkg

import "fmt"

/*
FileSizeFormatter é uma função responsável por fornecer um formatador
que padroniza o tamanho de um arquivo em medidas como MB, KB ,TB....
*/
func FileSizeFormatter(size int64) string {

	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case size >= TB:
		return fmt.Sprintf("%.2f TB", float64(size)/float64(TB))
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d size", size)
	}
}
