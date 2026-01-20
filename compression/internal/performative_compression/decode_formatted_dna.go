package performativecompression

import (
	"fmt"
	"io"
	"strings"
)

var dnaDecode = []byte{'A', 'C', 'G', 'T'}

type BitReader struct {
	data   []byte
	bitPos int
}

func (br *BitReader) ReadBits(n int) (uint8, error) {
	var v uint8
	for i := 0; i < n; i++ {
		bytePos := br.bitPos / 8
		if bytePos >= len(br.data) {
			return 0, io.EOF
		}
		bit := (br.data[bytePos] >> (7 - (br.bitPos % 8))) & 1
		v = (v << 1) | bit
		br.bitPos++
	}
	return v, nil
}

func (br *BitReader) RemainingBits() int {
	return len(br.data)*8 - br.bitPos
}

func Decode(data []byte) (string, error) {
	br := &BitReader{data: data}

	/* Captura do número de sequências */
	seqCount, err := br.ReadBits(8)
	if err != nil {
		return "", err
	}

	/* Captura das sequências puras */
	seqMap := make(map[uint8]string)

	for i := uint8(0); i < seqCount; i++ {
		var seq strings.Builder
		for j := 0; j < 5; j++ {
			v, _ := br.ReadBits(2)
			seq.WriteByte(dnaDecode[v])
		}
		seqMap[i] = seq.String()
	}

	/* payload */
	var out strings.Builder

	for br.RemainingBits() >= 4 { // evita ler padding
		mode, err := br.ReadBits(2)
		if err != nil {
			break
		}

		switch mode {

		// Convertendo os valores empacotados no modelo de DNA puro
		case 0b00:
			v, _ := br.ReadBits(2)
			out.WriteByte(dnaDecode[v])

		// Convertendo os valores empacotados no modelo de DNA com referências para as sequências
		case 0b01:
			class, _ := br.ReadBits(1)
			if class != 1 {
				return "", fmt.Errorf("classe inválida: %d", class)
			}

			refIdx, _ := br.ReadBits(3)
			seq, ok := seqMap[refIdx]
			if !ok {
				return "", fmt.Errorf("referência inexistente: %d", refIdx)
			}
			out.WriteString(seq)

		default:
			return "", fmt.Errorf("modo inválido: %b", mode)
		}
	}

	return out.String(), nil
}
