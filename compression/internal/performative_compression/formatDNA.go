package performativecompression

import "strings"

/*
formatDNA é uma função responsável por trocar os blocos de dna
que se repetem por um valor mapeado que irá representar esse bloco de agora em diante.
*/
// Sequência de DNA : ABCDABCDG
//
// Bloco com mais de uma recorrência : BCD -> I
//
// Retorno : AIAIG
func formatDNA(dna string) (string, []string, error) {
	buffer := dna
	dnaSequences := []string{}
	mapValues := []string{"I", "J", "P", "O"}

	for _, value := range mapValues {
		seq, err := analyzeDNASequency(buffer, 5)
		if err != nil || seq == "" {
			break // não há mais compressão possível
		}

		buffer = strings.ReplaceAll(buffer, seq, value)
		dnaSequences = append(dnaSequences, seq)
	}

	return buffer, dnaSequences, nil
}
