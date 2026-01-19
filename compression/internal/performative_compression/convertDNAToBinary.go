package performativecompression

/*
converDNAToBinary é uma função responsável por receber a string do dna previamente formatada e condensa-la em binário.
A função deve não só retornar o equivalente do dna em binário mas, também retornar os headers que contem
as sequências de dna que se repetem como também o número de sequências que vão ser salvas.
*/
func convertDNAToBinary(formattedDNA string, DNASequences []string) ([]byte, error) {
	var headers = make(map[string][]byte)

	// Adicionar ao header o metadata que contém o número de sequências salvas.
	var sequenceLen = []byte{
		byte(len(DNASequences)),
	}

	headers["sequenceLen"] = sequenceLen

	// Adicionar ao header os bytes que contém á sequência
	var sequenceBuffer byte
	var bitCount uint8

	var replacerSequence = map[rune]uint8{
		'A': 0b00,
		'C': 0b01,
		'G': 0b10,
		'T': 0b11,
	}

	for _, seq := range DNASequences {
		for _, sequenceRune := range seq {
			encoding, ok := replacerSequence[sequenceRune]

			if !ok {
				continue
			}

			sequenceBuffer = (sequenceBuffer << 2) | encoding
			bitCount += 2

			if bitCount == 8 {
				headers["sequence"] = append(headers["sequence"], sequenceBuffer)
				sequenceBuffer = 0
				bitCount = 0
			}
		}
	}

	// Adiciona ao header os dados do dna convertido para binario
	payload, errOnEncodeFormattedDna := encodeFormattedDNA(formattedDNA)

	if errOnEncodeFormattedDna != nil {
		return nil, errOnEncodeFormattedDna
	}

	var buffer []byte
	buffer = append(buffer, headers["sequenceLen"]...)
	buffer = append(buffer, headers["sequence"]...)
	buffer = append(buffer, payload...)

	return buffer, nil
}
