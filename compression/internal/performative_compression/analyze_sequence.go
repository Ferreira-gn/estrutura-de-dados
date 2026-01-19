package performativecompression

import "errors"

type resultItem struct {
	seq   string
	count int
}

/*
analyzeDNASequency é uma função responsável por analizar uma string
de valores presentes em uma representação do DNA e retornar a 
sequência com o maior número de recorrências.
*/
// String : AACTAATAACTAGTACTATG
//
// Sequência : AACT
//
// Número de recorrencias : 2
func analyzeDNASequency(dna string, windowSize int) (string, error) {
	
	if len(dna) < windowSize {
		return "", errors.New("sequência de DNA menor que 5 caracteres")
	}

	recurrencesOfSequence := make(map[string]int)
	order := make([]string, 0)

	// mapeamento de todas as sequências possíveis presente em um DNA
	for i := 0; i <= len(dna)-windowSize; i++ {
		seq := dna[i : i+windowSize]

		if containsControlChar(seq) {
			continue
		}

		if _, exists := recurrencesOfSequence[seq]; !exists {
			order = append(order, seq)
		}

		recurrencesOfSequence[seq]++
	}

	results := make([]resultItem, 0)

	for _, seq := range order {
		if recurrencesOfSequence[seq] > 1 {
			results = append(results, resultItem{
				seq:   seq,
				count: recurrencesOfSequence[seq],
			})
		}
	}

	if len(results) == 0 {
		return "", errors.New("nenhuma sequência válida de 5 caracteres se repete")
	}

	// Ordenação por maior recorrência
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].count > results[i].count {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	return results[0].seq, nil
}

/*
containsControlChar é uma função responsável por impedir que os valores 
I , J , O e P sejam usados como parâmetro para a formação de uma sequência válida.
 */
func containsControlChar(seq string) bool {
	for i := 0; i < len(seq); i++ {
		switch seq[i] {
		case 'I', 'J', 'O', 'P':
			return true
		}
	}
	return false
}
